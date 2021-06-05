package euler2

import "testing"

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []struct {
		desc           string
		actualResult   int
		expectedResult int
	}{
		{"21: Amicable Numbers", AmicableNumbers(10000), 31626},
		{"22: Name Scores", NameScores(), 871198282},
	}
	for _, tC := range testCases {
		if actualResult := tC.actualResult; actualResult != tC.expectedResult {
			t.Errorf("Failure. Test case: %v. Actual result: %v. Expected result: %v.", tC.desc, tC.actualResult, tC.expectedResult)
		}
	}
}
