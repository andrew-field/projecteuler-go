package numbertheory

import "testing"

func TestGetPrimeFactorisationANDGetPrimeFactorisationInSlice(t *testing.T) {
	testCases := []struct {
		input          uint
		expectedResult []uint
	}{
		{0, []uint{0}},
		{1, []uint{1}},
		{2, []uint{2}},
		{3, []uint{3}},
		{4, []uint{2, 2}},
		{5, []uint{5}},
		{6, []uint{2, 3}},
		{7, []uint{7}},
		{8, []uint{2, 2, 2}},
		{9, []uint{3, 3}},
		{10, []uint{2, 5}},
		{100, []uint{2, 2, 5, 5}},
		{101, []uint{101}},
		{1000, []uint{2, 2, 2, 5, 5, 5}},
		{4561356, []uint{2, 2, 3, 593, 641}},
		{600851475143, []uint{71, 839, 1471, 6857}},
	}
	for _, tC := range testCases {
		primeFactorChannel := GetPrimeFactorisation(tC.input)
		for _, expectedPrimeFactor := range tC.expectedResult {
			// All these tests should not encounter the problem where the prime factor channel has too few values. The reverse is checked below.
			if actualPrimeFactor := <-primeFactorChannel; expectedPrimeFactor != actualPrimeFactor {
				t.Errorf("GetPrimeFactorisation has failed. Input in test: %v. Expected prime: %v. Actual prime: %v.", tC.input, expectedPrimeFactor, actualPrimeFactor)
			}
		}

		// Check the prime factor channel does not have too many values.
		if _, unfinished := <-primeFactorChannel; unfinished {
			t.Errorf("GetPrimeFactorisation has failed. Input in test: %v. The prime factor channel has too many primes", tC.input)
		}

		actualResult := GetPrimeFactorisationInSlice(tC.input)
		if len(tC.expectedResult) != len(actualResult) {
			t.Errorf("GetPrimeFactorisationInSlice has failed. Input in test: %v. Expected result: %v. Actual result: %v.", tC.input, tC.expectedResult, actualResult)
		}
		for ind, expectedPrimeFactor := range tC.expectedResult {
			if actualPrimeFactor := actualResult[ind]; expectedPrimeFactor != actualPrimeFactor {
				t.Errorf("GetPrimeFactorisationInSlice has failed. Input in test: %v. Expected prime: %v. Actual prime: %v.", tC.input, expectedPrimeFactor, actualPrimeFactor)
			}
		}
	}
}

func TestLargestPrimeFactor(t *testing.T) {
	testCases := []struct {
		input          uint
		expectedResult uint
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 2},
		{5, 5},
		{6, 3},
		{7, 7},
		{8, 2},
		{9, 3},
		{10, 5},
		{100, 5},
		{101, 101},
		{1000, 5},
		{4561356, 641},
		{600851475143, 6857},
	}
	for _, tC := range testCases {
		if actualResult := LargestPrimeFactor(tC.input); tC.expectedResult != actualResult {
			t.Errorf("LargestPrimeFactor has failed. Input in test: %v. Expected largest prime factor: %v. Actual largest prime factor: %v.", tC.input, tC.expectedResult, actualResult)
		}
	}
}
