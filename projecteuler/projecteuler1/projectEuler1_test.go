package projecteuler1

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc           string
		function       func() int
		expectedResult int
	}{
		{"1: Multiples of 3 and 5 (Method 1)", MultiplesOf3And5MethodOne, 233168},
		{"1: Multiples of 3 and 5 (Method 2)", MultiplesOf3And5MethodTwo, 233168},
		{"2: Even Fibonacci Numbers", EvenFibonacciNumbers, 4613732},
		{"3: Largest Prime Factor", LargestPrimeFactor, 6857},
	}
	for _, tC := range testCases {
		actualResult := tC.function()
		if tC.expectedResult != actualResult {
			t.Errorf("Failure. Test case: %#v. Expected result: %#v. Actual result: %#v", tC.desc, tC.expectedResult, actualResult)
		}
	}
}
