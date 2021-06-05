package euler1

import (
	"github.com/andrew-field/maths"
)

// HighlyDivisibleTriangularNumber returns the value of the first triangle number to have over |n| divisors.
func HighlyDivisibleTriangularNumber(n int) int {
	n = maths.Abs(n)

	// Each triangular number.
	triangularNumber := 0

	// Calculate the number of divisors and see if it is greater than 500.
	sequenceStep := 1
	for numberOfDivisors := 0; numberOfDivisors < n; sequenceStep++ {
		triangularNumber += sequenceStep
		numberOfDivisors = maths.NumberOfDivisors(triangularNumber)
	}

	return triangularNumber
}
