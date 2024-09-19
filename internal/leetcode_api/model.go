package leetcode_api

type Submission struct {
	Count      int64  `json:"count"`
	Difficulty string `json:"difficulty"`
}

type SubmitStats struct {
	AcSubmissionNum []Submission `json:"acSubmissionNum"`
}

type UserProfileData struct {
	RealName        string       `json:"realName"`
	UserSlug        string       `json:"userSlug"`
	Rank            float64      `json:"rank"`
	AllProblemCount []Submission `json:"allProblemCount"`
}

type Profile struct {
	UserSlug string  `json:"userSlug"`
	Rank     float64 `json:"ranking"`
}

type MatchedUser struct {
	Username    string      `json:"username"`
	Profile     Profile     `json:"profile"`
	SubmitStats SubmitStats `json:"submitStats"`
}
