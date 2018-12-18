package numbertheory

import (
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestGetNumberOfDigitsOfAnInt(t *testing.T) {
	testCases := []struct {
		number                 int
		expectedNumberOfDigits int
	}{
		{0, 1},
		{1, 1},
		{9, 1},
		{10, 2},
		{99, 2},
		{100, 3},
		{500, 3},
		{math.MaxInt32, 10},
		{math.MaxInt64, 19},
	}
	for ind, tC := range testCases {
		t.Run(strconv.Itoa(ind), func(t *testing.T) {
			for index := 0; index < 2; index++ {
				// For each number, try the negative as well.
				tC.number *= -1
				t.Logf("tC.number: %#+v\n", tC.number)
				if actualNumberOfDigits := GetNumberOfDigitsOfAnInt(tC.number); actualNumberOfDigits != tC.expectedNumberOfDigits {
					t.Errorf("Number in test: %v. Expected number of digits: %v. Actual number of digits: %v.", tC.number, tC.expectedNumberOfDigits, actualNumberOfDigits)
				}
			}
		})
	}

	// Test random numbers.
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for index := 0; index < 1000; index++ {
		// This will give a random number n such that 0 <= n < 18.
		powerOfTen := int(randomGenerator.Float64() * 18)
		lowerBound := math.Pow10(powerOfTen)
		// This will give a random number n such that lowerBound <= n < lowerBound*10 where lowerBound is always a multiple of 10.
		// The expected number of digits is powerOfTen + 1.
		testNumber := int(randomGenerator.Float64()*(9*lowerBound) + lowerBound)
		t.Run(strconv.Itoa(testNumber), func(t *testing.T) {
			for index := 0; index < 2; index++ {
				// For each number, try the negative as well.
				testNumber *= -1
				t.Logf("testNumber: %#+v\n", testNumber)
				if actualNumberOfDigits := GetNumberOfDigitsOfAnInt(testNumber); actualNumberOfDigits != powerOfTen+1 {
					t.Errorf("Number in test: %v. Expected number of digits: %v. Actual number of digits: %v.", testNumber, powerOfTen, actualNumberOfDigits)
				}
			}
		})
	}
}

func TestGetNumberOfDigitsOfABigInt(t *testing.T) {
	testCases := []struct {
		number                 *big.Int
		expectedNumberOfDigits int
	}{
		{big.NewInt(0), 1},
		{big.NewInt(1), 1},
		{big.NewInt(9), 1},
		{big.NewInt(10), 2},
		{big.NewInt(99), 2},
		{big.NewInt(100), 3},
		{big.NewInt(500), 3},
		{big.NewInt(math.MaxInt32), 10},
		{big.NewInt(math.MaxInt64), 19},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(20), nil), 20},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(35), nil), 35},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(99), nil), 99},
		{big.NewInt(0).Exp(big.NewInt(10), big.NewInt(100), nil), 100},
	}
	for ind, tC := range testCases {
		t.Run(strconv.Itoa(ind), func(t *testing.T) {
			for index := 0; index < 2; index++ {
				// For each number, try the negative as well.
				tC.number.Mul(tC.number, big.NewInt(-1))
				t.Logf("tC.number: %#+v\n", tC.number)
				if actualNumberOfDigits := GetNumberOfDigitsOfABigInt(tC.number); actualNumberOfDigits != tC.expectedNumberOfDigits {
					t.Errorf("Number in test: %v. Expected number of digits: %v. Actual number of digits: %v.", tC.number, tC.expectedNumberOfDigits, actualNumberOfDigits)
				}
			}
		})
	}

	// Test random numbers.
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for index := 0; index < 1000; index++ {
		// This will give a random number n such that 0 <= n < 100.
		powerOfTen := int64(randomGenerator.Float64() * 100)
		lowerBound := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(powerOfTen), nil)
		// This will give a random number n such that lowerBound <= n < lowerBound*10 where lowerBound is always a multiple of 10.
		// The expected number of digits is powerOfTen + 1.
		nine := big.NewInt(9)
		rand := big.NewInt(int64(randomGenerator.Float64()))
		testNumber := lowerBound.Add(rand.Mul(rand, nine.Mul(nine, lowerBound)), lowerBound)
		t.Run(testNumber.String(), func(t *testing.T) {
			for index := 0; index < 2; index++ {
				// For each number, try the negative as well.
				testNumber.Mul(testNumber, big.NewInt(-1))
				t.Logf("testNumber: %#+v\n", testNumber)
				if actualNumberOfDigits := GetNumberOfDigitsOfABigInt(testNumber); actualNumberOfDigits != int(powerOfTen+1) {
					t.Errorf("Number in test: %v. Expected number of digits: %v. Actual number of digits: %v.", testNumber, powerOfTen, actualNumberOfDigits)
				}
			}
		})
	}
}
