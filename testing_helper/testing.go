package testing_helper

import "testing"

type TestCase struct {
	Desc                         string
	ActualResult, ExpectedResult int
}

func CheckResults(testCases []TestCase, t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.Desc, func(t *testing.T) {
			if tC.ActualResult != tC.ExpectedResult {
				t.Errorf("Actual result: %d. Expected result: %d.", tC.ActualResult, tC.ExpectedResult)
			}
		})
	}
}
