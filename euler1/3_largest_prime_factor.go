package euler1

import "github.com/andrew-field/maths/v2"

// largestPrimeFactor returns the largest prime factor of |x|.
func largestPrimeFactor(x int) int {
	largestPrimeFactor := 0

	for v := range maths.PrimeFactorisation(x) {
		largestPrimeFactor = v.Value
	}

	return largestPrimeFactor
}
