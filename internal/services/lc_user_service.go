package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cmd/internal/db"
	"cmd/internal/leetcode_api"
	"cmd/internal/storage/dbs"
	"database/sql"

	v1 "cmd/internal/templates/v1"
)

type LcUserService struct {
	repository *db.Repository
}

func NewLcUserService(repository *db.Repository) *LcUserService {
	return &LcUserService{repository: repository}
}

func (s *LcUserService) GetOrCreate(ctx context.Context, userSlug string) (*v1.LcUserData, error) {
	dbUserData, err := s.repository.Queries().UserGetStatsBySlug(ctx, userSlug)
	deprecatedStats := dbUserData.UpdatedAt.Before(time.Now().Add(-15 * time.Minute))

	if errors.Is(err, sql.ErrNoRows) {
		matchedUser, err := leetcode_api.MatchedUserMapToUserProfile(userSlug)
		userData := v1.NewLcUserDataFromReq(*matchedUser)
		if err != nil {
			return userData, fmt.Errorf("GetOrCreate: error getting user stats in user_service from lc_api: %w", err)
		}

		// Inserting user stats to lc_stats table
		err = s.InsertUserStats(ctx, userData)
		if err != nil {
			return userData, fmt.Errorf("GetOrCreate: error inserting new user_stats in user_service: %w", err)
		}

		return userData, nil
	} else if deprecatedStats {
		matchedUser, err := leetcode_api.MatchedUserMapToUserProfile(userSlug)
		userData := v1.NewLcUserDataFromReq(*matchedUser)
		if err != nil {
			return userData, fmt.Errorf("GetOrCreate: error getting user stats in user_service from lc_api: %w", err)
		}

		err = s.UpdateUserStats(ctx, userData)
		if err != nil {
			return userData, fmt.Errorf("GetOrCreate: error updating user stats in user_service from lc_api: %w", err)
		}

		return userData, nil
	} else if err != nil {
		return nil, fmt.Errorf("GetOrCreate: if err != nil && err != sql.ErrNoRows in user_service: %w", err)
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
