package numbertheory

import (
	"math"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func TestGetNumberOfDigitsOfAnInt(t *testing.T) {
	testCases := []struct {
		input          int
		expectedResult int
	}{
		{0, 1},
		{1, 1},
		{9, 1},
		{10, 2},
		{99, 2},
		{100, 3},
		{500, 3},
		{4563198, 7},
		{math.MaxInt32, 10},
		{math.MaxInt64 - 1, 19},
	}
	for _, tC := range testCases {
		for index := 0; index < 2; index++ {
			// For each number, try the negative as well.
			tC.input *= -1
			if actualNumberOfDigits := GetNumberOfDigitsOfAnInt(tC.input); actualNumberOfDigits != tC.expectedResult {
				t.Errorf("Input in test: %v. Expected number of digits: %v. Actual number of digits: %v.", tC.input, tC.expectedResult, actualNumberOfDigits)
			}
		}
	}

	// Test random numbers.
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for index := 0; index < 1000; index++ {
		// This will give a random number n such that 0 <= n < 18.
		powerOfTen := int(randomGenerator.Float64() * 18)
		// lowerBound is of the form 10^x.
		lowerBound := math.Pow10(powerOfTen)
		// This will give a random number n such that lowerBound <= n < lowerBound*10.
		input := int(randomGenerator.Float64()*(9*lowerBound) + lowerBound)
		// The expected number of digits is powerOfTen + 1.
		expectedResult := powerOfTen + 1
		for index := 0; index < 2; index++ {
			// For each number, try the negative as well.
			input *= -1
			if actualNumberOfDigits := GetNumberOfDigitsOfAnInt(input); actualNumberOfDigits != expectedResult {
				t.Errorf("Input in test: %v. Expected number of digits: %v. Actual number of digits: %v.", input, expectedResult, actualNumberOfDigits)
			}
		}
	}
}

func TestGetNumberOfDigitsOfABigInt(t *testing.T) {
	testCases := []struct {
		input          *big.Int
		expectedResult int
	}{
		{big.NewInt(0), 1},
		{big.NewInt(1), 1},
		{big.NewInt(9), 1},
		{big.NewInt(10), 2},
		{big.NewInt(99), 2},
		{big.NewInt(100), 3},
		{big.NewInt(500), 3},
		{big.NewInt(4563198), 7},
		{big.NewInt(math.MaxInt32), 10},
		{big.NewInt(math.MaxInt64), 19},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(20), nil), 21},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(35), nil), 36},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(98), nil), 99},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(99), nil), 100},
	}
	for _, tC := range testCases {
		for index := 0; index < 2; index++ {
			// For each number, try the negative as well.
			tC.input.Mul(tC.input, big.NewInt(-1))
			if actualNumberOfDigits := GetNumberOfDigitsOfABigInt(tC.input); actualNumberOfDigits != tC.expectedResult {
				t.Errorf("Input in test: %v. Expected number of digits: %v. Actual number of digits: %v.", tC.input, tC.expectedResult, actualNumberOfDigits)
			}
		}
	}

	// Test random numbers.
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for index := 0; index < 1000; index++ {
		// This will give a random number n such that 0 <= n < 100.
		powerOfTen := int64(randomGenerator.Float64() * 100)
		// lowerBound is of the form 10^x.
		lowerBound := big.NewFloat(0).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(powerOfTen), nil))
		// This will give a random number n such that lowerBound <= n < lowerBound*10,
		nine := big.NewFloat(9)
		rand := big.NewFloat(randomGenerator.Float64())
		input, _ := lowerBound.Add(rand.Mul(rand, nine.Mul(nine, lowerBound)), lowerBound).Int(nil)
		// The expected number of digits is powerOfTen + 1.
		expectedResult := powerOfTen + 1
		for index := 0; index < 2; index++ {
			// For each number, try the negative as well.
			input.Mul(input, big.NewInt(-1))
			if actualNumberOfDigits := GetNumberOfDigitsOfABigInt(input); actualNumberOfDigits != int(expectedResult) {
				t.Errorf("Input in test: %v. Expected number of digits: %v. Actual number of digits: %v.", input, expectedResult, actualNumberOfDigits)
			}
		}
	}
}

func TestGetDigitsOfAnIntANDGetDigitsOfAnIntInSlice(t *testing.T) {
	testCases := []struct {
		input          int
		expectedDigits []int
	}{
		{0, []int{0}},
		{1, []int{1}},
		{9, []int{9}},
		{10, []int{0, 1}},
		{99, []int{9, 9}},
		{100, []int{0, 0, 1}},
		{500, []int{0, 0, 5}},
		{4563198, []int{8, 9, 1, 3, 6, 5, 4}},
		{math.MaxInt32, []int{7, 4, 6, 3, 8, 4, 7, 4, 1, 2}},
		{math.MaxInt64, []int{7, 0, 8, 5, 7, 7, 4, 5, 8, 6, 3, 0, 2, 7, 3, 3, 2, 2, 9}},
	}
	for _, tC := range testCases {
		for index := 0; index < 2; index++ {
			// For each number, try the negative as well.
			tC.input *= -1
			digitChannel := GetDigitsOfAnInt(tC.input)
			digitSlice := GetDigitsOfAnIntInSlice(tC.input)
			index := 0
			for actualDigit := range digitChannel {
				if actualDigit != tC.expectedDigits[index] {
					t.Errorf("GetDigitsOfAnInt has failed. Input in test: %v. Expected digit: %v. Actual digit: %v.", tC.input, tC.expectedDigits[index], actualDigit)
				}
				if actualDigitFromSlice := digitSlice[len(digitSlice)-1-index]; actualDigitFromSlice != tC.expectedDigits[index] {
					t.Errorf("GetDigitsOfAnIntInSlice has failed. Input in test: %v. Expected digit: %v. Actual digit: %v.", tC.input, tC.expectedDigits[index], actualDigitFromSlice)
				}
				index++
			}
		}
	}
}

func TestGetDigitsOfABigIntANDGetDigitsOfABigNumberInSlice(t *testing.T) {
	testCases := []struct {
		input          *big.Int
		expectedDigits []int
	}{
		{big.NewInt(0), []int{0}},
		{big.NewInt(1), []int{1}},
		{big.NewInt(9), []int{9}},
		{big.NewInt(10), []int{0, 1}},
		{big.NewInt(99), []int{9, 9}},
		{big.NewInt(100), []int{0, 0, 1}},
		{big.NewInt(500), []int{0, 0, 5}},
		{big.NewInt(4563198), []int{8, 9, 1, 3, 6, 5, 4}},
		{big.NewInt(math.MaxInt32), []int{7, 4, 6, 3, 8, 4, 7, 4, 1, 2}},
		{big.NewInt(math.MaxInt64), []int{7, 0, 8, 5, 7, 7, 4, 5, 8, 6, 3, 0, 2, 7, 3, 3, 2, 2, 9}},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(20), nil), []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{big.NewInt(0).Exp(big.NewInt(2), big.NewInt(100), nil), []int{6, 7, 3, 5, 0, 2, 3, 0, 7, 6, 9, 4, 1, 0, 4, 9, 2, 2, 8, 2, 2, 0, 0, 6, 0, 5, 6, 7, 6, 2, 1}},
	}
	for _, tC := range testCases {
		for ind := 0; ind < 2; ind++ {
			// For each number, try the negative as well.
			tC.input.Mul(tC.input, big.NewInt(-1))

			digitChannel := GetDigitsOfABigInt(tC.input)
			digitSlice := GetDigitsOfABigNumberInSlice(tC.input)
			index := 0
			for actualDigit := range digitChannel {
				if actualDigit != tC.expectedDigits[index] {
					t.Errorf("GetDigitsOfABigInt has failed. Input in test: %v. Expected digit: %v. Actual digit: %v.", tC.input, tC.expectedDigits[index], actualDigit)
				}
				if actualDigitFromSlice := digitSlice[len(digitSlice)-1-index]; actualDigitFromSlice != tC.expectedDigits[index] {
					t.Errorf("GetDigitsOfABigNumberInSlice has failed. Input in test: %v. Expected digit: %v. Actual digit: %v.", tC.input, tC.expectedDigits[index], actualDigitFromSlice)
				}
				index++
			}
		}
	}
}
