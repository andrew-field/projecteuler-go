package euler2

import "testing"

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []struct {
		desc           string
		actualResult   int
		expectedResult int
	}{
		{"21: Amicable Numbers", amicableNumbers(10000), 31626},
		{"22: Name Scores", nameScores(), 871198282},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actualResult := tC.actualResult; actualResult != tC.expectedResult {
				t.Errorf("Actual result: %d. Expected result: %d.", tC.actualResult, tC.expectedResult)
			}
		})
	}
}
