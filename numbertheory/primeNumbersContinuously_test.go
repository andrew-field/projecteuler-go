package numbertheory

import (
	"testing"
)

func TestGetPrimeNumbersContinuously(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult []uint
	}{
		{1, []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}},
		{10, []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}},
		{100, []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}},
		{1000, []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}},
		{10000, []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}},
	}
	for _, tC := range testCases {
		primeChannel, doneChannel := GetPrimeNumbersContinuously(tC.input)
		for _, expectedPrime := range tC.expectedResult {
			// All these tests should not encounter the problem where the prime channel has too few values.
			if actualPrime := <-primeChannel; expectedPrime != actualPrime {
				t.Errorf("GetPrimeNumbersContinuously has failed. Input in test: %v. Expected prime: %v. Actual prime: %v.", tC.input, expectedPrime, actualPrime)
			}
		}
		doneChannel <- true
	}

	primeChannelContinuous, doneChannel := GetPrimeNumbersContinuously(1000)
	primeChannelCeiling := GetAllPrimeNumbersBelowCeiling(1000000)

	for ceilingPrime := range primeChannelCeiling {
		if continuousPrime := <-primeChannelContinuous; uint(ceilingPrime) != continuousPrime {
			t.Errorf("GetPrimeNumbersContinuously has failed to match GetPrimeNumbersContinuously. Ceiling prime: %v. Continuous prime: %v.", ceilingPrime, continuousPrime)
		}
	}
	doneChannel <- true

	// Testing that the code panics if there is a bad input.
	CheckGetPrimeNumbersContinuouslyPanic(t, GetPrimeNumbersContinuously, 0)
	CheckGetPrimeNumbersContinuouslyPanic(t, GetPrimeNumbersContinuously, -10)
}

func CheckGetPrimeNumbersContinuouslyPanic(t *testing.T, function func(int) (chan uint, chan bool), inputForFunction int) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetPrimeNumbersContinuously has failed. The code did not panic.")
		}
	}()
	function(inputForFunction)
}
