package v1

import "cmd/internal/leetcode_api"

const (
	EasyMaxValue   int64 = 819
	MediumMaxValue int64 = 1710
	HardMaxValue   int64 = 732
	BarWidthValue  int64 = 240
)

type LcUserData struct {
	UserSlug    string
	Username    string
	EasyCount   int64
	MediumCount int64
	HardCount   int64
	TotalCount  int64
	Rank        float64
}

type BarsWidth struct {
	EasyWidth   float64
	MediumWidth float64
	HardWidth   float64
}

type VisitsStats struct {
	DailyVisits   int64
	WeeklyVisits  int64
	MonthlyVisits int64
	TotalVisits   int64
}

func NewLcUserData(username string, userSlug string, easyCount int64, mediumCount int64, hardCount int64, totalCount int64) *LcUserData {
	var result = &LcUserData{
		UserSlug:    userSlug,
		Username:    username,
		EasyCount:   easyCount,
		MediumCount: mediumCount,
		HardCount:   hardCount,
		TotalCount:  totalCount,
	}
	return result
}

func NewVisitsStats(DailyVisits int64, WeeklyVisits int64, MonthlyVisits int64, TotalVisits int64) *VisitsStats {
	return &VisitsStats{
		DailyVisits:   DailyVisits,
		WeeklyVisits:  WeeklyVisits,
		MonthlyVisits: MonthlyVisits,
		TotalVisits:   TotalVisits,
	}

}

func NewLcUserDataFromReq(profileData leetcode_api.UserProfileData) *LcUserData {
	var (
		easyCount   int64
		mediumCount int64
		hardCount   int64
		totalCount  int64
	)
	username := profileData.Username
	rank := profileData.Rank
	userSlug := profileData.UserSlug

	for _, problem := range profileData.AllProblemCount {
		switch problem.Difficulty {
		case "Easy":
			totalCount += problem.Count
			easyCount += problem.Count
		case "Medium":
			totalCount += problem.Count
			mediumCount += problem.Count
		case "Hard":
			totalCount += problem.Count
			hardCount += problem.Count
		}
	}

	var result = &LcUserData{Username: username, UserSlug: userSlug, EasyCount: easyCount, MediumCount: mediumCount, HardCount: hardCount, TotalCount: totalCount, Rank: rank}
	return result
}
