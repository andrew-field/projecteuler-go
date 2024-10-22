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
		if tC.actualResult != tC.expectedResult {
			t.Errorf("Failure. Test case: %v. Actual result: %v. Expected result: %v.", tC.desc, tC.actualResult, tC.expectedResult)
		}
	}
}
