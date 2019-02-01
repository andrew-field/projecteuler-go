package projecteuler

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc           string
		expectedResult int
	}{
		{"1: Multiples of 3 and 5", 233168},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actualResult := MultiplesOf3And5MethodOne()
			if tC.expectedResult != actualResult {
				t.Errorf("Failure. Test case: %#v. Expected result: %#v. Actual result: %#v", tC.desc, tC.expectedResult, actualResult)
			}
		})
	}
}
