package euler1

import (
	"github.com/andrew-field/maths/v2"
)

// highlyDivisibleTriangularNumber returns the value of the first triangle number to have over |n| divisors.
func highlyDivisibleTriangularNumber(n int) int {
	n, err := maths.Abs(n)
	if err != nil {
		panic(err)
	}

	// Each triangular number.
	triangularNumber := 0

	// Calculate the number of divisors and see if it is greater than |n|.
	sequenceStep := 1
	for numberOfDivisors := 0; numberOfDivisors < n; sequenceStep++ {
		triangularNumber += sequenceStep
		numberOfDivisors = maths.NumberOfDivisorsBruteForce(triangularNumber)
	}

	return triangularNumber
}
