package euler4

import "testing"

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []struct {
		desc           string
		actualResult   int
		expectedResult int
	}{
		{"67: Maximum Path Sum Two", maximumPathSumTwo(), 7273},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.actualResult != tC.expectedResult {
				t.Errorf("Actual result: %d. Expected result: %d.", tC.actualResult, tC.expectedResult)
			}
		})
	}
}
