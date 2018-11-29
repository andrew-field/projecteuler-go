package numbertheory

import (
	"math"
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
				tC.number *= -1
				t.Logf("tC.number: %#+v\n", tC.number)
				if actualNumberOfDigits := GetNumberOfDigitsOfAnInt(tC.number); actualNumberOfDigits != tC.expectedNumberOfDigits {
					t.Errorf("Number in test: %v. Expected number of digits: %v. Actual number of digits: %v.", tC.number, tC.expectedNumberOfDigits, actualNumberOfDigits)
				}
			}
		})
	}

	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for powerOfTen := 0; powerOfTen < 20; powerOfTen++ {
		limit := math.Pow10(powerOfTen)
		testNumber := int(randomGenerator.Float64() + limit)
		t.Run(strconv.Itoa(powerOfTen), func(t *testing.T) {
			for index := 0; index < 2; index++ {
				testNumber *= -1
				t.Logf("testNumber: %#+v\n", testNumber)
				if actualNumberOfDigits := GetNumberOfDigitsOfAnInt(testNumber); actualNumberOfDigits != powerOfTen+1 {
					t.Errorf("Number in test: %v. Expected number of digits: %v. Actual number of digits: %v.", testNumber, powerOfTen, actualNumberOfDigits)
				}
			}
		})
	}
}

/* func TestGetNumberOfDigitsOfAFloat(t *testing.T) {
	testCases := []struct {
		number                 float64
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
				tC.number *= -1
				if actualNumberOfDigits := GetNumberOfDigitsOfAFloat(tC.number); actualNumberOfDigits != tC.expectedNumberOfDigits {
					t.Errorf("Number in test: %v. Expected number of digits: %v. Actual number of digits: %v.", tC.number, tC.expectedNumberOfDigits, actualNumberOfDigits)
				}
			}
		})
	}
} */
