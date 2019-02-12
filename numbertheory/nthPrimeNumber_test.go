package numbertheory

import "testing"

func TestGetNthPrimeNumber(t *testing.T) {
	testCases := []struct {
		input          uint
		expectedResult uint
	}{
		{1, 2},
		{2, 3},
		{3, 5},
		{4, 7},
		{5, 11},
		{6, 13},
		{7, 17},
		{8, 19},
		{9, 23},
		{10, 29},
		{100, 541},
		{101, 547},
		{1000, 7919},
		{10001, 104743},
	}
	for _, tC := range testCases {
		if actualResult := GetNthPrimeNumber(tC.input); tC.expectedResult != actualResult {
			t.Errorf("GetNthPrimeNumber has failed. Input in test: %v. Expected prime: %v. Actual prime: %v.", tC.input, tC.expectedResult, actualResult)
		}
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetNthPrimeNumber has failed. The code did not panic.")
		}
	}()
	GetNthPrimeNumber(0)
}
