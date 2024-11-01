// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package dbs

import (
	"time"
)

type LcStat struct {
	UserSlug      string
	Username      string
	EasySubmits   int64
	MediumSubmits int64
	HardSubmits   int64
	TotalSubmits  int64
	Rank          int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ProfileHourlyVisitsStat struct {
	UserSlug string
	Time     time.Time
	Count    int64
}
