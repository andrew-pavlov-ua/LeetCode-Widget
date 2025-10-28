package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"cmd/internal/db"
	"cmd/internal/env"
	"cmd/internal/leetcode_api"
	"cmd/internal/storage/dbs"

	"database/sql"

	v1 "cmd/internal/templates/v1"

	"github.com/redis/go-redis/v9"
)

type LcUserService struct {
	mutex      sync.Mutex
	repository *db.Repository
	rdb        *redis.Client
}

func NewLcUserService(repository *db.Repository, rdb *redis.Client) *LcUserService {
	return &LcUserService{repository: repository, rdb: rdb}
}

func (s *LcUserService) GetOrCreate(ctx context.Context, userSlug string) (*v1.LcUserData, error) {
	var userData *v1.LcUserData
	var err error

	if env.Must("USE_REDIS") == "y" {
		userData, err = s.RedisGetOrCreate(ctx, userSlug)
	} else {
		userData, err = s.SQLGetOrCreate(ctx, userSlug)
	}

	if err != nil {
		return nil, err
	}

	return userData, nil
}

// Using Redis for stats storing gives 240x faster response than default postgres according to tests
// Redis: Time (mean ± σ): 2.6 ms ±   0.2 ms 1000 runs
// Postgres: Time (mean ± σ): 625.6 ms ± 234.6 ms 1000 runs
func (s *LcUserService) RedisGetOrCreate(ctx context.Context, userSlug string) (*v1.LcUserData, error) {
	var userData v1.LcUserData
	// formatting user key like user:murasame_
	redis_user_id := fmt.Sprintf("user:%s", userSlug)
	err := s.rdb.HGetAll(ctx, redis_user_id).Scan(&userData)
	if err == redis.Nil || userData.Username == "" {
		// No user data with userSlug, getting it from Leetcode API
		expirationTime := 15 * time.Minute
		// API request
		matchedUser, err := leetcode_api.MatchedUserMapToUserProfile(userSlug)
		userData := v1.NewLcUserDataFromReq(*matchedUser)
		if err != nil {
			return userData, fmt.Errorf("RedisGetOrCreate: error getting user stats in user_service from lc_api: %w", err)
		}

		// Preparing hash fields to store in Redis
		matchedUserHash := []string{
			"user_slug", userData.UserSlug,
			"username", userData.Username,
			"easy_count", fmt.Sprint(userData.EasyCount),
			"medium_count", fmt.Sprint(userData.MediumCount),
			"hard_count", fmt.Sprint(userData.HardCount),
			"total_count", fmt.Sprint(userData.TotalCount),
			"rank", strconv.FormatFloat(userData.Rank, 'f', -1, 64),
		}
		s.rdb.HSet(ctx, redis_user_id, matchedUserHash)
		s.rdb.Expire(ctx, redis_user_id, expirationTime)
		return userData, nil
	} else if err != nil {
		return nil, fmt.Errorf("RedisGetOrCreate: if err != nil && err != sql.ErrNoRows in user_service: %w", err)
	}

	return &userData, nil
}

func (s *LcUserService) SQLGetOrCreate(ctx context.Context, userSlug string) (*v1.LcUserData, error) {
	log.Print("GetOrCreate using postgres")
	dbUserData, err := s.repository.Queries().UserGetStatsBySlug(ctx, userSlug)
	deprecatedStats := dbUserData.UpdatedAt.Before(time.Now().Add(-15 * time.Minute))

	if errors.Is(err, sql.ErrNoRows) {
		// No user data with userSlug, getting it from Leetcode API
		matchedUser, err := leetcode_api.MatchedUserMapToUserProfile(userSlug)
		userData := v1.NewLcUserDataFromReq(*matchedUser)
		if err != nil {
			return userData, fmt.Errorf("SQLGetOrCreate: error getting user stats in user_service from lc_api: %w", err)
		}

		// Inserting user stats to lc_stats table
		err = s.InsertUserStats(ctx, userData)
		if err != nil {
			return userData, fmt.Errorf("SQLGetOrCreate: error inserting new user_stats in user_service: %w", err)
		}

		return userData, nil
	} else if deprecatedStats {
		matchedUser, err := leetcode_api.MatchedUserMapToUserProfile(userSlug)
		userData := v1.NewLcUserDataFromReq(*matchedUser)
		if err != nil {
			return userData, fmt.Errorf("SQLGetOrCreate: error getting user stats in user_service from lc_api: %w", err)
		}

		err = s.UpdateUserStats(ctx, userData)
		if err != nil {
			return userData, fmt.Errorf("SQLGetOrCreate: error updating user stats in user_service from lc_api: %w", err)
		}

		return userData, nil
	} else if err != nil {
		return nil, fmt.Errorf("SQLGetOrCreate: if err != nil && err != sql.ErrNoRows in user_service: %w", err)
	}

	userData := &v1.LcUserData{
		UserSlug:    userSlug,
		Username:    dbUserData.Username,
		EasyCount:   dbUserData.EasySubmits,
		MediumCount: dbUserData.MediumSubmits,
		HardCount:   dbUserData.HardSubmits,
		TotalCount:  dbUserData.TotalSubmits,
		Rank:        float64(dbUserData.Rank)}

	return userData, nil
}

func (s *LcUserService) UpdateUserStats(ctx context.Context, userData *v1.LcUserData) error {
	now := time.Now().UTC()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	err := s.repository.Queries().UpdateLcStats(ctx,
		dbs.UpdateLcStatsParams{
			EasySubmits:   userData.EasyCount,
			MediumSubmits: userData.MediumCount,
			HardSubmits:   userData.HardCount,
			TotalSubmits:  userData.TotalCount,
			UpdatedAt:     now})

	return err
}

func (s *LcUserService) InsertUserStats(ctx context.Context, userData *v1.LcUserData) error {
	now := time.Now().UTC()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	err := s.repository.Queries().InsertStatsInfo(ctx, dbs.InsertStatsInfoParams{
		Rank:          int64(userData.Rank),
		UserSlug:      userData.UserSlug,
		Username:      userData.Username,
		EasySubmits:   userData.EasyCount,
		MediumSubmits: userData.MediumCount,
		HardSubmits:   userData.HardCount,
		TotalSubmits:  userData.TotalCount,
		CreatedAt:     now,
		UpdatedAt:     now})

	return err
}

func (s *LcUserService) GetUserIdBySlug(ctx context.Context, userSlug string) (int64, error) {
	return s.repository.Queries().GetIdBySlug(ctx, userSlug)
}
