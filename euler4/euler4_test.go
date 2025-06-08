package euler4

import (
	"testing"

	"github.com/andrew-field/projecteuler-go/testing_helper"
)

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []testing_helper.TestCase{
		{Desc: "67: Maximum Path Sum Two", ActualResult: maximumPathSumTwo(), ExpectedResult: 7273},
	}

	testing_helper.CheckResults(testCases, t)
}
