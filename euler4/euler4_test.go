package euler4

import (
	"testing"

	"github.com/andrew-field/projecteuler-go/test"
)

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []test.Case{
		{Desc: "67: Maximum Path Sum Two", ActualResult: maximumPathSumTwo(), ExpectedResult: 7273},
	}

	test.CheckResults(testCases, t)
}
