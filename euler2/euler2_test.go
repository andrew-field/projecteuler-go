package euler2

import (
	"testing"

	"github.com/andrew-field/projecteuler-go/euler1"
)

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []euler1.TestCase{
		{"21: Amicable Numbers", amicableNumbers(10000), 31626},
		{"22: Name Scores", nameScores(), 871198282},
	}

	euler1.CheckResults(testCases, t)
}
