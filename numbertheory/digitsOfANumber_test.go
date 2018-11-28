package numbertheory

import (
	"strconv"
	"testing"
)

func TestGetNumberOfDigitsOfANumber(t *testing.T) {
	testCases := []struct {
		number                 float64
		expectedNumberOfDigits float64
	}{
		{-500, 3},
		{-100, 3},
		{-99, 2},
		{-10, 2},
		{-9, 1},
		{0, 1},
		{9, 1},
		{10, 2},
		{99, 2},
		{100, 3},
		{500, 3},
	}
	for ind, tC := range testCases {
		t.Run(strconv.Itoa(ind), func(t *testing.T) {
			if actualNumberOfDigits := GetNumberOfDigitsOfANumber(tC.number); actualNumberOfDigits != tC.expectedNumberOfDigits {
				t.Errorf("Actual number of digits: %v. Expected number of digits: %v. Number in test: %v", actualNumberOfDigits, tC.expectedNumberOfDigits, tC.number)
			}
		})
	}
}
