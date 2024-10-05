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

func getQueryQntyQuestions() string {
	return `{ allQuestionsCount { 
			difficulty 
			count 
			}
		}`
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

func MatchedUserMapToUserProfile(username string) *UserProfileData {
	matchedUser, err := getUserProfile(username)
	if err != nil || matchedUser["matchedUser"] == nil {
		fmt.Printf("Error getting matched user: %v", err)
		return &UserProfileData{
			Username: username + " user doesn't exist",
			UserSlug: "",
			Rank:     0,
			AllProblemCount: []Submission{
				{Count: 0, Difficulty: "Easy"},
				{Count: 0, Difficulty: "Medium"},
				{Count: 0, Difficulty: "Hard"},
			},
		}
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

	return &profileData
}

//doesn't work :(

//func GetUserProfile(username string) (*UserProfileData, error) {
//	var (
//		matchedUser MatchedUser
//		requestUser UserProfileData
//	)
//
//	client := graphql.NewClient("https://leetcode.com/graphql")
//	query := getQueryUserData()
//	req := graphql.NewRequest(query)
//	req.Var("username", username)
//	ctx := context.Background()
//
//	// Debugging: Log the raw JSON response
//	var rawResponse map[string]interface{}
//	err := client.Run(ctx, req, &rawResponse)
//	if err != nil {
//		return nil, fmt.Errorf("error making GraphQL request: %v", err)
//	}
//	fmt.Printf("Raw GraphQL response: %+v\n", rawResponse)
//
//	err = client.Run(ctx, req, &matchedUser)
//	if err != nil {
//		return nil, fmt.Errorf("error making GraphQL request: %v", err)
//	}
//
//	fmt.Printf("GraphQL response: %+v\n", matchedUser)
//
//	requestUser.Username = username
//	requestUser.AllProblemCount = matchedUser.SubmitStats.AcSubmissionNum
//	requestUser.Rank = matchedUser.Profile.Rank
//	requestUser.UserSlug = matchedUser.Profile.UserSlug
//
//	return &requestUser, nil
//}
