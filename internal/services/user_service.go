package services

import (
	"cmd/internal/db"
	"cmd/internal/leetcode_api"
	"cmd/internal/storage/dbs"
	v1 "cmd/internal/templates/v1"
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
)

type UserService struct {
	repository *db.Repository
}

func NewUserService(repository *db.Repository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) Upsert(ctx context.Context, userSlug string) int64 {
	id, err := s.repository.Queries().UserGetBySocialProviderId(ctx, userSlug)

	// If there's no users in db with id we need, creating user
	if errors.Is(err, sql.ErrNoRows) {
		matchedUser := leetcode_api.MatchedUserMapToUserProfile(userSlug)
		userData := v1.NewLcUserDataFromReq(*matchedUser)
		log.Println("userData", userData)

		//creating new user in db (table users)
		id, err := s.repository.Queries().UserNewAndParse(ctx, dbs.UserNewAndParseParams{
			Username:             userData.Username,
			SocialProviderUserID: userData.UserSlug})
		if err != nil {
			log.Printf("Database error 39: %s\n", err)

			return 0
		}

		// Inserting user stats to lc_stats table
		err = s.InsertUserStats(ctx, userData, id)
		if err != nil {
			log.Printf("Database err 45: %s\n", err)
		}
		return id
	} else if err != nil {
		log.Printf("Database err 56: %s\n", err)
	}
	return id
}

func (s *UserService) GetByStatsById(ctx context.Context, userId int64) (*v1.LcUserData, error) {
	now := time.Now().UTC()
	var (
		userProfileData *v1.LcUserData
		err             error
	)
	//Getting userStats by id from db
	userStatsByIDRow, err := s.repository.Queries().UserGetStatsByID(ctx, userId)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("No user stats row: %s\n", err)
		return nil, err
	} else if err != nil {
		log.Printf("Database err 65: %s\n", err)
		return nil, err
	}

	// If user stats were changed more than 15mins ago, updating it
	if userStatsByIDRow.UpdatedAt.UTC().Before(now.Add(-15 * time.Minute)) {
		userProfileData := *(leetcode_api.MatchedUserMapToUserProfile(userStatsByIDRow.Userslug))
		lcData := v1.NewLcUserDataFromReq(userProfileData)
		err = s.UpdateUserStats(ctx, lcData, userId)
		log.Println("getting info from LC, difference: ", userStatsByIDRow.UpdatedAt.UTC().Sub(now).Minutes())
	}
	userProfileData = &v1.LcUserData{
		Username:    userStatsByIDRow.Username,
		UserSlug:    userStatsByIDRow.Userslug,
		Rank:        float64(userStatsByIDRow.Rank),
		EasyCount:   userStatsByIDRow.EasySubmits.Int64,
		MediumCount: userStatsByIDRow.MediumSubmits.Int64,
		HardCount:   userStatsByIDRow.HardSubmits.Int64,
		TotalCount:  userStatsByIDRow.TotalSubmits.Int64,
	}

	log.Println("Parsed user: ", userProfileData.UserSlug)
	return userProfileData, err
}

func (s *UserService) UpdateUserStats(ctx context.Context, userData *v1.LcUserData, userId int64) error {
	now := time.Now().UTC()

	err := s.repository.Queries().UpdateLcStats(ctx,
		dbs.UpdateLcStatsParams{
			EasySubmits: sql.NullInt64{
				Int64: userData.EasyCount,
				Valid: true},
			MediumSubmits: sql.NullInt64{
				Int64: userData.MediumCount,
				Valid: true},
			HardSubmits: sql.NullInt64{
				Int64: userData.HardCount,
				Valid: true},
			TotalSubmits: sql.NullInt64{
				Int64: userData.TotalCount,
				Valid: true},
			UpdatedAt: now,
			UserID:    userId})

	return err
}

func (s *UserService) InsertUserStats(ctx context.Context, userData *v1.LcUserData, userId int64) error {
	now := time.Now().UTC()

	err := s.repository.Queries().InsertStatsInfo(ctx, dbs.InsertStatsInfoParams{
		UserID: userId,
		Rank:   int64(userData.Rank),
		EasySubmits: sql.NullInt64{
			Int64: userData.EasyCount,
			Valid: true},
		MediumSubmits: sql.NullInt64{
			Int64: userData.MediumCount,
			Valid: true},
		HardSubmits: sql.NullInt64{
			Int64: userData.HardCount,
			Valid: true},
		TotalSubmits: sql.NullInt64{
			Int64: userData.TotalCount,
			Valid: true},
		CreatedAt: now,
		UpdatedAt: now})

	return err
}
