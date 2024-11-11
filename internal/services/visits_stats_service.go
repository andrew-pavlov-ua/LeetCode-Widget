package services

import (
	"context"
	"fmt"
	"time"

	"cmd/internal/db"
	"cmd/internal/storage/dbs"

	v1 "cmd/internal/templates/v1"
)

type VisitsStatsService struct {
	repository *db.Repository
}

func NewVisistsStatsService(repository *db.Repository) *VisitsStatsService {
	return &VisitsStatsService{repository: repository}
}

func (s *VisitsStatsService) Upsert(ctx context.Context, userSlug string) error {
	now := time.Now().UTC()

	err := s.repository.Queries().ProfileHourlyVisitsStatsUpsert(ctx, dbs.ProfileHourlyVisitsStatsUpsertParams{
		UserSlug: userSlug,
		Time:     now,
		Count:    1,
	})
	return err
}

func (s *VisitsStatsService) TotalCount(ctx context.Context, userSlug string) (int64, error) {
	return s.repository.Queries().TotalCount(ctx, userSlug)
}

func (s *VisitsStatsService) InsertCount(ctx context.Context, userSlug string) error {
	now := time.Now().UTC().Truncate(time.Hour)

	err := s.repository.Queries().ProfileHourlyVisitsStatsUpsert(ctx, dbs.ProfileHourlyVisitsStatsUpsertParams{
		UserSlug: userSlug,
		Time:     now,
		Count:    1,
	})
	return err
}

func (s *VisitsStatsService) GetFullStatsCount(ctx context.Context, userSlug string) (v1.VisitsStats, error) {
	var (
		totalVisitCount int64
		visitStats      v1.VisitsStats
		now             = time.Now().UTC()
	)

	rawViewsStats, err := s.repository.Queries().ProfileHourlyViewsStats(ctx, dbs.ProfileHourlyViewsStatsParams{
		Day:      now.AddDate(0, 0, -1),
		Week:     now.AddDate(0, 0, -7),
		Month:    now.AddDate(0, -1, 0),
		UserSlug: userSlug})

	if err != nil {
		return visitStats, fmt.Errorf("GetFullStatsCount: error with count requesting: %w", err)
	}

	totalVisitCount, err = s.repository.Queries().TotalCount(ctx, userSlug)
	if err != nil {
		return visitStats, fmt.Errorf("GetFullStatsCount: error with total count requesting: %w", err)
	}

	visitStats = *v1.NewVisitsStats(rawViewsStats.DayCount, rawViewsStats.WeekCount, rawViewsStats.MonthCount, totalVisitCount)
	return visitStats, nil
}
