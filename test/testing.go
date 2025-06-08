package test

import "testing"

type Case struct {
	Desc                         string
	ActualResult, ExpectedResult int
}

func CheckResults(testCases []Case, t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.Desc, func(t *testing.T) {
			if tC.ActualResult != tC.ExpectedResult {
				t.Errorf("Actual result: %d. Expected result: %d.", tC.ActualResult, tC.ExpectedResult)
			}
		})
	}
}
