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

func GetUserProfile(username string) (*UserProfileData, error) {
	var (
		matchedUser MatchedUser
		requestUser UserProfileData
	)

	client := graphql.NewClient("https://leetcode.com/graphql")
	query := getQueryUserData()
	req := graphql.NewRequest(query)
	req.Var("username", username)
	ctx := context.Background()

	// Try executing the request and handle any errors
	err := client.Run(ctx, req, &matchedUser)
	if err != nil {
		return nil, fmt.Errorf("error making GraphQL request: %v", err)
	}

	// Debugging: Print the response for inspection
	fmt.Printf("GraphQL response: %+v\n", matchedUser)

	requestUser.Username = username
	requestUser.AllProblemCount = matchedUser.SubmitStats.AcSubmissionNum
	requestUser.Rank = matchedUser.Profile.Rank
	requestUser.UserSlug = matchedUser.Profile.UserSlug

	return &requestUser, nil
}
