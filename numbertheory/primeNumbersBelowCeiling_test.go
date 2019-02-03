package numbertheory

import "testing"

func TestGetAllPrimeNumbersBelowCeiling(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult []int
	}{
		{2, []int{2}},
		{7, []int{2, 3, 5, 7}},
		{8, []int{2, 3, 5, 7}},
		{9, []int{2, 3, 5, 7}},
		{10, []int{2, 3, 5, 7}},
		{11, []int{2, 3, 5, 7, 11}},
		{16, []int{2, 3, 5, 7, 11, 13}},
		{17, []int{2, 3, 5, 7, 11, 13, 17}},
		{18, []int{2, 3, 5, 7, 11, 13, 17}},
		{19, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{21, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{22, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{97, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
		{98, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
		{99, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
		{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
		{101, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}},
	}
	for _, tC := range testCases {
		primeChannel := GetAllPrimeNumbersBelowCeiling(tC.input)
		for _, expectedPrime := range tC.expectedResult {
			// All these tests should not encounter the problem where the prime channel has too few values. The reverse is checked below.
			if actualPrime := <-primeChannel; expectedPrime != actualPrime {
				t.Errorf("GetAllPrimeNumbersBelowCeiling has failed. Input in test: %#v. Expected prime: %v. Actual prime: %v.", tC.input, expectedPrime, actualPrime)
			}
		}

		// Check the prime channel does not have too many values.
		if _, unfinished := <-primeChannel; unfinished {
			t.Errorf("GetAllPrimeNumbersBelowCeiling has failed. Input in test: %#v. The prime channel has too many primes", tC.input)
		}
	}

	// Testing that the code panics if there is a bad input.
	CheckGetAllPrimeNumbersBelowCeilingPanic(t, GetAllPrimeNumbersBelowCeiling, 1)
	CheckGetAllPrimeNumbersBelowCeilingPanic(t, GetAllPrimeNumbersBelowCeiling, -10)
}

func CheckGetAllPrimeNumbersBelowCeilingPanic(t *testing.T, function func(int) chan int, inputForFunction int) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetAllPrimeNumbersBelowCeiling has failed. The code did not panic.")
		}
	}()
	function(inputForFunction)
}
