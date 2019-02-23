package numbertheory

import "testing"

func TestGetMaximumPathSumOfPyramidUsingMaximumSlotsANDGetMaximumPathSumOfPyramidUsingRecursiveFunction(t *testing.T) {
	testCases := []struct {
		input          [][]float64
		expectedResult float64
	}{
		{[][]float64{{0}}, 0},
		{[][]float64{{1}}, 1},
		{[][]float64{{1}, {2, 3}}, 4},
		{[][]float64{{1}, {2, 3}, {4, 5, 6}}, 10},
		{[][]float64{{75}, {95, 64}, {17, 47, 82}, {18, 35, 87, 10}}, 308},
		{[][]float64{{75}, {95, 64}, {17, 47, 82}, {18, 35, 87, 10}, {24, 12, 54, 38, 20}, {46, 35, 42, 64, 21, 45}}, 426},
	}
	type functionToTest struct {
		name     string
		function func([][]float64) float64
	}
	for _, tC := range testCases {
		functions := []functionToTest{{"GetMaximumPathSumOfPyramidUsingMaximumSlots", GetMaximumPathSumOfPyramidUsingMaximumSlots},
			{"GetMaximumPathSumOfPyramidUsingRecursiveFunction", GetMaximumPathSumOfPyramidUsingRecursiveFunction}}
		for _, function := range functions {
			if actualResult := function.function(tC.input); tC.expectedResult != actualResult {
				t.Errorf("%v has failed. Input in test: %v. Expected result: %v. Actual result: %v.", function.name, tC.input, tC.expectedResult, actualResult)
			}
		}
	}

	// Testing that the code panics if there is a bad input.
	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingMaximumSlots, "GetMaximumPathSumOfPyramidUsingMaximumSlots", nil)
	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingMaximumSlots, "GetMaximumPathSumOfPyramidUsingMaximumSlots", [][]float64{{1, 2}})
	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingMaximumSlots, "GetMaximumPathSumOfPyramidUsingMaximumSlots", [][]float64{{1}, {2, 3, 4}})
	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingMaximumSlots, "GetMaximumPathSumOfPyramidUsingMaximumSlots", make([][]float64, 0))
	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingRecursiveFunction, "GetMaximumPathSumOfPyramidUsingRecursiveFunction", nil)
	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingRecursiveFunction, "GetMaximumPathSumOfPyramidUsingRecursiveFunction", [][]float64{{1, 2}})
	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingRecursiveFunction, "GetMaximumPathSumOfPyramidUsingRecursiveFunction", [][]float64{{1}, {2, 3, 4}})
	CheckGetMaximumPathSumOfPyramidPanics(t, GetMaximumPathSumOfPyramidUsingRecursiveFunction, "GetMaximumPathSumOfPyramidUsingRecursiveFunction", make([][]float64, 0))

}

func CheckGetMaximumPathSumOfPyramidPanics(t *testing.T, function func([][]float64) float64, functionName string, inputForFunction [][]float64) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%v has failed. The code did not panic.", functionName)
		}
	}()
	function(inputForFunction)
}
