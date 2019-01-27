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
		input                  int
		expectedNumberOfDigits int
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
		{math.MaxInt64, 19},
	}
	for _, tC := range testCases {
		for index := 0; index < 2; index++ {
			// For each number, try the negative as well.
			tC.input *= -1
			if actualNumberOfDigits := GetNumberOfDigitsOfAnInt(tC.input); actualNumberOfDigits != tC.expectedNumberOfDigits {
				t.Errorf("Number in test: %#v. Expected number of digits: %#v. Actual number of digits: %#v.", tC.input, tC.expectedNumberOfDigits, actualNumberOfDigits)
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
		expectedNumberOfDigits := powerOfTen + 1
		for index := 0; index < 2; index++ {
			// For each number, try the negative as well.
			input *= -1
			if actualNumberOfDigits := GetNumberOfDigitsOfAnInt(input); actualNumberOfDigits != expectedNumberOfDigits {
				t.Errorf("Number in test: %#v. Expected number of digits: %#v. Actual number of digits: %#v.", input, expectedNumberOfDigits, actualNumberOfDigits)
			}
		}
	}
}

func TestGetNumberOfDigitsOfABigInt(t *testing.T) {
	testCases := []struct {
		input                  *big.Int
		expectedNumberOfDigits int
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
			if actualNumberOfDigits := GetNumberOfDigitsOfABigInt(tC.input); actualNumberOfDigits != tC.expectedNumberOfDigits {
				t.Errorf("Number in test: %#v. Expected number of digits: %#v. Actual number of digits: %#v.", tC.input, tC.expectedNumberOfDigits, actualNumberOfDigits)
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
		expectedNumberOfDigits := powerOfTen + 1
		for index := 0; index < 2; index++ {
			// For each number, try the negative as well.
			input.Mul(input, big.NewInt(-1))
			if actualNumberOfDigits := GetNumberOfDigitsOfABigInt(input); actualNumberOfDigits != int(expectedNumberOfDigits) {
				t.Errorf("Number in test: %#v. Expected number of digits: %#v. Actual number of digits: %#v.", input, expectedNumberOfDigits, actualNumberOfDigits)
			}
		}
	}
}
