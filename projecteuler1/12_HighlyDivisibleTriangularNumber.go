package projecteuler1

import (
	"github.com/andrew-field/testing_go/numbertheory"
)

// HighlyDivisibleTriangularNumber returns the value of the first triangle number to have over five hundred divisors.
func HighlyDivisibleTriangularNumber() int {

	// Each triangular number.
	var triangularNumber uint = 1

	// Calculate the number of divisors and see if it is greater than 500.
	var sequenceStep uint = 2
	for divisors := 1; divisors < 501; sequenceStep++ {
		triangularNumber += sequenceStep
		divisors = numbertheory.GetNumberOfDivisors(triangularNumber)
	}

	return int(triangularNumber)
}
