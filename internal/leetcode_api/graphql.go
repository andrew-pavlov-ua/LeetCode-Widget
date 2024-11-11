package leetcode_api

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

func getQueryUserData() string {
	return `query getUserProfile($username: String!) {
  				matchedUser(username: $username) {
    				username
    				profile {
					realName
      				ranking
      				userSlug
    				}
    				submitStats {
      				acSubmissionNum {
        				difficulty
        				count
      				}
    				}
  				}
				}
`
}

func getUserProfile(username string) (map[string]interface{}, error) {
	client := graphql.NewClient("https://leetcode.com/graphql")
	query := getQueryUserData()
	req := graphql.NewRequest(query)
	req.Var("username", username)
	ctx := context.Background()

	var response map[string]interface{}
	err := client.Run(ctx, req, &response)
	if err != nil {
		return nil, fmt.Errorf("error making GraphQL request: %v", err)
	}

	return response, nil
}

func MatchedUserMapToUserProfile(userSlug string) (*UserProfileData, error) {
	matchedUser, err := getUserProfile(userSlug)
	if err != nil {
		return &UserProfileData{
			Username: "User with lcId " + userSlug + " doesn't exist",
			UserSlug: "",
			Rank:     0,
			AllProblemCount: []Submission{
				{Count: 0, Difficulty: "Easy"},
				{Count: 0, Difficulty: "Medium"},
				{Count: 0, Difficulty: "Hard"},
			},
		}, fmt.Errorf("MatchedUserMapToUserProfile: requesting user from ;lc api: %w", err)
	}

	_username := matchedUser["matchedUser"].(map[string]interface{})["profile"].(map[string]interface{})["realName"].(string)
	_userSlug := matchedUser["matchedUser"].(map[string]interface{})["profile"].(map[string]interface{})["userSlug"].(string)
	_rank := matchedUser["matchedUser"].(map[string]interface{})["profile"].(map[string]interface{})["ranking"].(float64)

	profileData := UserProfileData{
		Username: _username,
		UserSlug: _userSlug,
		Rank:     _rank,
	}

	for _, item := range matchedUser["matchedUser"].(map[string]interface{})["submitStats"].(map[string]interface{})["acSubmissionNum"].([]interface{}) {
		sub := item.(map[string]interface{})
		profileData.AllProblemCount = append(profileData.AllProblemCount, Submission{
			Count:      int64(sub["count"].(float64)),
			Difficulty: sub["difficulty"].(string),
		})
	}

	return &profileData, nil
}
