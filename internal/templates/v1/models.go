package v1

import "cmd/internal/leetcode_api"

const (
	EasyMaxValue   int64 = 819
	MediumMaxValue int64 = 1710
	HardMaxValue   int64 = 732
	BarWidthValue  int64 = 240
)

type LcUserData struct {
<<<<<<< HEAD
	UserSlug    string
=======
>>>>>>> 78514ac7933200ad2012e1f91d3e06f9481c63c3
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

<<<<<<< HEAD
func NewLcUserData(userSlug string, easyCount int64, mediumCount int64, hardCount int64, totalCount int64) *LcUserData {
	var result = &LcUserData{
		UserSlug:    userSlug,
=======
func NewLcUserData(username string, easyCount int64, mediumCount int64, hardCount int64, totalCount int64) *LcUserData {
	var result = &LcUserData{
		Username:    username,
>>>>>>> 78514ac7933200ad2012e1f91d3e06f9481c63c3
		EasyCount:   easyCount,
		MediumCount: mediumCount,
		HardCount:   hardCount,
		TotalCount:  totalCount,
	}
	return result
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
<<<<<<< HEAD
	userSlug := profileData.UserSlug
=======
>>>>>>> 78514ac7933200ad2012e1f91d3e06f9481c63c3

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

<<<<<<< HEAD
	var result = &LcUserData{Username: username, UserSlug: userSlug, EasyCount: easyCount, MediumCount: mediumCount, HardCount: hardCount, TotalCount: totalCount, Rank: rank}
=======
	var result = &LcUserData{Username: username, EasyCount: easyCount, MediumCount: mediumCount, HardCount: hardCount, TotalCount: totalCount, Rank: rank}
>>>>>>> 78514ac7933200ad2012e1f91d3e06f9481c63c3
	return result
}
