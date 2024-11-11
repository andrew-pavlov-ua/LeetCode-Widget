package testing

import (
	. "cmd/internal/leetcode_api"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchedUserMapToUserProfile(t *testing.T) {
	t.Logf("Started testing lc api")
	userSlug := "MURASAME_"
	allSubmissions := &Submission{Count: 27, Difficulty: "All"}
	ezSubmissions := &Submission{Count: 24, Difficulty: "Easy"}
	medSubmissions := &Submission{Count: 3, Difficulty: "Medium"}
	hardSubmissions := &Submission{Count: 0, Difficulty: "Hard"}

	actualUser, err := MatchedUserMapToUserProfile(userSlug)

	if err != nil {
		fmt.Printf("TestMatchedUserMapToUserProfile: err testing: %e", err)
		assert.Failf(t, "TestMatchedUserMapToUserProfile: err", "err testing: %e", err)
	}

	expectedUser := UserProfileData{
		Username:        "Andrew",
		UserSlug:        userSlug,
		Rank:            actualUser.Rank, // getting rank from actualUser 'cause it's changing often
		AllProblemCount: []Submission{*allSubmissions, *ezSubmissions, *medSubmissions, *hardSubmissions},
	}

	assert.Equal(t, expectedUser, *actualUser, "Parsed user data should be same as simulated")
}
