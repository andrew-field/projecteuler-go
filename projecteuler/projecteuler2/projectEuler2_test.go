package projecteuler2

import "testing"

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []struct {
		desc           string
		function       func() int
		expectedResult int
	}{
		{"21: Amicable Numbers", AmicableNumbers, 31626},
		{"22: Name Scores", NameScores, 871198282},
	}
	for _, tC := range testCases {
		if actualResult := tC.function(); tC.expectedResult != actualResult {
			t.Errorf("Failure. Test case: %v. Expected result: %v. Actual result: %v", tC.desc, tC.expectedResult, actualResult)
		}
	}
}
