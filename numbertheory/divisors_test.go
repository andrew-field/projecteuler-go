package numbertheory

import (
	"sort"
	"testing"
)

func TestGetNumberOfDivisors(t *testing.T) {
	testCases := []struct {
		input          uint
		expectedResult int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 3},
		{5, 2},
		{6, 4},
		{7, 2},
		{8, 4},
		{9, 3},
		{10, 4},
		{100, 9},
		{500, 12},
		{45664, 12},
		{7931265, 32},
	}
	for _, tC := range testCases {
		if actualResult := GetNumberOfDivisors(tC.input); actualResult != tC.expectedResult {
			t.Errorf("GetNumberOfDivisors has failed. Input in test: %v. Expected number of divisors: %v. Actual number of divisors: %v.", tC.input, tC.expectedResult, actualResult)
		}
	}
}

func TestGetDivisorsOfANumberANDGetDivisorsOfANumberInASlice(t *testing.T) {
	testCases := []struct {
		input          uint
		expectedResult []uint
	}{
		{1, []uint{1}},
		{2, []uint{1, 2}},
		{3, []uint{1, 3}},
		{4, []uint{1, 2, 4}},
		{5, []uint{1, 5}},
		{6, []uint{1, 2, 3, 6}},
		{7, []uint{1, 7}},
		{8, []uint{1, 2, 4, 8}},
		{9, []uint{1, 3, 9}},
		{10, []uint{1, 2, 5, 10}},
		{100, []uint{1, 2, 4, 5, 10, 20, 25, 50, 100}},
		{500, []uint{1, 2, 4, 5, 10, 20, 25, 50, 100, 125, 250, 500}},
		{45664, []uint{1, 2, 4, 8, 16, 32, 1427, 2854, 5708, 11416, 22832, 45664}},
		{7931265, []uint{1, 3, 5, 15, 17, 51, 85, 255, 19, 57, 95, 285, 323, 969, 1615, 4845, 1637, 4911, 8185, 24555, 27829, 83487, 139145, 417435, 31103, 93309, 155515, 466545, 528751, 1586253, 2643755, 7931265}},
	}
	for _, tC := range testCases {
		// GetDivisorsOfANumber
		divisorChannel := GetDivisorsOfANumber(tC.input)
		// All these tests should not encounter the problem where the divisor channel has too few values. The reverse is checked below.
		for _, expectedResult := range tC.expectedResult {
			if actualResult := <-divisorChannel; actualResult != expectedResult {
				t.Errorf("GetDivisorsOfANumber has failed. Input in test: %v. Expected divisor: %v. Actual divisor: %v.", tC.input, expectedResult, actualResult)
			}
		}

		// Check the divisor channel does not have too many values.
		if _, unfinished := <-divisorChannel; unfinished {
			t.Errorf("GetDivisorsOfANumber has failed. Input in test: %v. The divisor channel has too many values", tC.input)
		}

		// GetDivisorsOfANumberInASlice
		sort.Slice(tC.expectedResult, func(i, j int) bool { return tC.expectedResult[i] < tC.expectedResult[j] })
		actualResult := GetDivisorsOfANumberInASlice(tC.input)
		if len(tC.expectedResult) != len(actualResult) {
			t.Errorf("GetDivisorsOfANumberInASlice has failed. Input in test: %v. Expected result: %v. Actual result: %v.", tC.input, tC.expectedResult, actualResult)
		}
		for ind, expectedDivisor := range tC.expectedResult {
			if actualDivisor := actualResult[ind]; expectedDivisor != actualDivisor {
				t.Errorf("GetDivisorsOfANumberInASlice has failed. Input in test: %v. Expected divisor: %v. Actual divisor: %v.", tC.input, expectedDivisor, actualDivisor)
			}
		}
	}

	// Testing that the code panics if there is a bad input.
	CheckGetDivisorsOfANumber(t, GetDivisorsOfANumber, "GetDivisorsOfANumber", 0)
	CheckGetDivisorsOfANumberInASlice(t, GetDivisorsOfANumberInASlice, "GetDivisorsOfANumberInASlice", 0)
}

func TestGetDivisorsOfANumberBruteForceANDGetDivisorsOfANumberInASliceBruteForce(t *testing.T) {
	testCases := []struct {
		input          uint
		expectedResult []uint
	}{
		{1, []uint{1}},
		{2, []uint{1, 2}},
		{3, []uint{1, 3}},
		{4, []uint{1, 4, 2}},
		{5, []uint{1, 5}},
		{6, []uint{1, 6, 2, 3}},
		{7, []uint{1, 7}},
		{8, []uint{1, 8, 2, 4}},
		{9, []uint{1, 9, 3}},
		{10, []uint{1, 10, 2, 5}},
		{100, []uint{1, 100, 2, 50, 4, 25, 5, 20, 10}},
		{500, []uint{1, 500, 2, 250, 4, 125, 5, 100, 10, 50, 20, 25}},
		{45664, []uint{1, 45664, 2, 22832, 4, 11416, 8, 5708, 16, 2854, 32, 1427}},
		{7931265, []uint{1, 7931265, 3, 2643755, 5, 1586253, 15, 528751, 17, 466545, 19, 417435, 51, 155515, 57, 139145, 85, 93309, 95, 83487, 255, 31103, 285, 27829, 323, 24555, 969, 8185, 1615, 4911, 1637, 4845}},
	}
	for _, tC := range testCases {
		// GetDivisorsOfANumberBruteForce
		divisorChannel := GetDivisorsOfANumberBruteForce(tC.input)
		// All these tests should not encounter the problem where the divisor channel has too few values. The reverse is checked below.
		for _, expectedResult := range tC.expectedResult {
			if actualResult := <-divisorChannel; actualResult != expectedResult {
				t.Errorf("GetDivisorsOfANumberBruteForce has failed. Input in test: %v. Expected divisor: %v. Actual divisor: %v.", tC.input, expectedResult, actualResult)
			}
		}

		// Check the divisor channel does not have too many values.
		if _, unfinished := <-divisorChannel; unfinished {
			t.Errorf("GetDivisorsOfANumberBruteForce has failed. Input in test: %v. The divisor channel has too many values", tC.input)
		}

		// GetDivisorsOfANumberInASliceBruteForce
		sort.Slice(tC.expectedResult, func(i, j int) bool { return tC.expectedResult[i] < tC.expectedResult[j] })
		actualResult := GetDivisorsOfANumberInASliceBruteForce(tC.input)
		if len(tC.expectedResult) != len(actualResult) {
			t.Errorf("GetDivisorsOfANumberInASliceBruteForce has failed. Input in test: %v. Expected result: %v. Actual result: %v.", tC.input, tC.expectedResult, actualResult)
		}
		for ind, expectedDivisor := range tC.expectedResult {
			if actualDivisor := actualResult[ind]; expectedDivisor != actualDivisor {
				t.Errorf("GetDivisorsOfANumberInASliceBruteForce has failed. Input in test: %v. Expected divisor: %v. Actual divisor: %v.", tC.input, expectedDivisor, actualDivisor)
			}
		}
	}

	// Testing that the code panics if there is a bad input.
	CheckGetDivisorsOfANumber(t, GetDivisorsOfANumberBruteForce, "GetDivisorsOfANumberBruteForce", 0)
	CheckGetDivisorsOfANumberInASlice(t, GetDivisorsOfANumberInASliceBruteForce, "GetDivisorsOfANumberInASliceBruteForce", 0)
}

func CheckGetDivisorsOfANumber(t *testing.T, function func(uint) chan uint, functionName string, inputForFunction uint) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%v has failed. The code did not panic.", functionName)
		}
	}()
	function(inputForFunction)
}

func CheckGetDivisorsOfANumberInASlice(t *testing.T, function func(uint) []uint, functionName string, inputForFunction uint) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%v has failed. The code did not panic.", functionName)
		}
	}()
	function(inputForFunction)
}
