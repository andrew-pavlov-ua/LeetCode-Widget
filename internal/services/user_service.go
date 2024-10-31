package services

import (
	"cmd/internal/db"
	"cmd/internal/leetcode_api"
	"cmd/internal/storage/dbs"
	v1 "cmd/internal/templates/v1"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
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
		matchedUser := leetcode_api.MatchedUserMapToUserProfile(userSlug)
		userData := v1.NewLcUserDataFromReq(*matchedUser)
		log.Println("userData", userData)

		// Inserting user stats to lc_stats table
		fmt.Println("Inserting user stats to lc_stats table")
		err = s.InsertUserStats(ctx, userData)
		if err != nil {
			log.Println("err 37: ", err)
		}
		return userData, err
	} else if err != nil && err != sql.ErrNoRows {
		fmt.Println("err41 : ", err)
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
		LcUserID:      userData.UserSlug,
		Username:      userData.Username,
		EasySubmits:   userData.EasyCount,
		MediumSubmits: userData.MediumCount,
		HardSubmits:   userData.HardCount,
		TotalSubmits:  userData.TotalCount,
		CreatedAt:     now,
		UpdatedAt:     now})

	return err
}
