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

type UserService struct {
	repository *db.Repository
}

func NewUserService(repository *db.Repository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) Upsert(ctx context.Context, userSlug string) (*v1.LcUserData, error) {
	dbUserData, err := s.repository.Queries().UserGetStatsBySlug(ctx, userSlug)

	if errors.Is(err, sql.ErrNoRows) {
		matchedUser, err := leetcode_api.MatchedUserMapToUserProfile(userSlug)
		userData := v1.NewLcUserDataFromReq(*matchedUser)
		if err != nil {
			return userData, fmt.Errorf("Upsert: error getting user stats in user_service from lc_api: %w", err)
		}

		// Inserting user stats to lc_stats table
		err = s.InsertUserStats(ctx, userData)
		if err != nil {
			return userData, fmt.Errorf("Upsert: error inserting new user_stats in user_service: %w", err)
		}
		return userData, nil
	} else if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("Upsert: if err != nil && err != sql.ErrNoRows in user_service: %w", err)
	}
	userData := &v1.LcUserData{
		UserSlug:    userSlug,
		Username:    dbUserData.Username,
		EasyCount:   dbUserData.EasySubmits,
		MediumCount: dbUserData.MediumSubmits,
		HardCount:   dbUserData.HardSubmits,
		TotalCount:  dbUserData.TotalSubmits,
		Rank:        float64(dbUserData.Rank)}

	if dbUserData.UpdatedAt.Before(time.Now().UTC().Add(-15 * time.Minute)) {
		s.UpdateUserStats(ctx, userData)
	}

	return userData, nil
}

func (s *UserService) UpdateUserStats(ctx context.Context, userData *v1.LcUserData) error {
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

func (s *UserService) InsertUserStats(ctx context.Context, userData *v1.LcUserData) error {
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
