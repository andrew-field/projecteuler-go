package euler4

import (
	"testing"

	"github.com/andrew-field/projecteuler-go/euler1"
)

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []euler1.TestCase{
		{"67: Maximum Path Sum Two", maximumPathSumTwo(), 7273},
	}

	euler1.CheckResults(testCases, t)
}
