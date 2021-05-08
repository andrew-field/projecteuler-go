package projecteuler4

import "testing"

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []struct {
		desc           string
		function       func() int
		expectedResult int
	}{
		{"67: Maximum Path Sum Two", MaximumPathSumTwo, 7273},
	}
	for _, tC := range testCases {
		if actualResult := tC.function(); tC.expectedResult != actualResult {
			t.Errorf("Failure. Test case: %v. Expected result: %v. Actual result: %v", tC.desc, tC.expectedResult, actualResult)
		}
	}
}
