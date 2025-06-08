package euler2

import (
	"testing"

	"github.com/andrew-field/projecteuler-go/testing_helper"
)

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []testing_helper.TestCase{
		{Desc: "21: Amicable Numbers", ActualResult: amicableNumbers(10000), ExpectedResult: 31626},
		{Desc: "22: Name Scores", ActualResult: nameScores(), ExpectedResult: 871198282},
	}

	testing_helper.CheckResults(testCases, t)
}
