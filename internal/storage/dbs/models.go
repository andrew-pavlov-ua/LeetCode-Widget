// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package dbs

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type SocialProvider string

const (
	SocialProviderLeetcode SocialProvider = "leetcode"
)

func (e *SocialProvider) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = SocialProvider(s)
	case string:
		*e = SocialProvider(s)
	default:
		return fmt.Errorf("unsupported scan type for SocialProvider: %T", src)
	}
	return nil
}

type NullSocialProvider struct {
	SocialProvider SocialProvider
	Valid          bool // Valid is true if SocialProvider is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullSocialProvider) Scan(value interface{}) error {
	if value == nil {
		ns.SocialProvider, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.SocialProvider.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullSocialProvider) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.SocialProvider), nil
}

type LcStat struct {
	UserID        int64
	Username      string
	EasySubmits   sql.NullInt64
	MediumSubmits sql.NullInt64
	HardSubmits   sql.NullInt64
	TotalSubmits  sql.NullInt64
	Rank          int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type User struct {
	ID       int64
	LcUserID string
}
