package leetcode_api

type Submission struct {
	Count      int64  `json:"count"`
	Difficulty string `json:"difficulty"`
}

type SubmitStats struct {
	AcSubmissionNum []Submission `json:"acSubmissionNum"`
}

type UserProfileData struct {
	Username        string       `json:"username"`
	UserSlug        string       `json:"user_slug"`
	Rank            int64        `json:"rank"`
	Lvl             int64        `json:"lvl"`
	AllProblemCount []Submission `json:"all_question"`
}

type Profile struct {
	UserSlug string `json:"user_slug"`
	Rank     int64  `json:"ranking"`
	Lvl      int64  `json:"level"`
}

type MatchedUser struct {
	Username    string      `json:"username"`
	SubmitStats SubmitStats `json:"submit_stats"`
	Profile     Profile     `json:"profile"`
}
