package euler2

import (
	"testing"

	"github.com/andrew-field/projecteuler-go/test"
)

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []test.TestCase{
		{Desc: "21: Amicable Numbers", ActualResult: amicableNumbers(10000), ExpectedResult: 31626},
		{Desc: "22: Name Scores", ActualResult: nameScores(), ExpectedResult: 871198282},
	}

	test.CheckResults(testCases, t)
}
