package numbertheory

import "testing"

func TestLowestCommonMultiple(t *testing.T) {
	testCases := []struct {
		input          []uint
		expectedResult uint
	}{
		{[]uint{0}, 0},
		{[]uint{1}, 1},
		{[]uint{0, 1}, 0},
		{[]uint{1, 2}, 2},
		{[]uint{2, 3, 4}, 12},
		{[]uint{6, 10}, 30},
		{[]uint{5, 10, 65}, 130},
		{[]uint{75, 225, 330, 450}, 4950},
		{[]uint{13, 89, 456}, 527592},
	}
	for _, tC := range testCases {
		if actualResult := LowestCommonMultiple(tC.input...); tC.expectedResult != actualResult {
			t.Errorf("LowestCommonMultiple has failed. Input in test: %v. Expected result: %v. Actual result: %v.", tC.input, tC.expectedResult, actualResult)
		}
	}
}
