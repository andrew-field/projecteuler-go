package euler1

import "github.com/andrew-field/maths"

// LargestPrimeFactor returns the largest prime factor of |x|.
// Does not handle math.MinInt64
func LargestPrimeFactor(x int) int {

	largestPrimeFactor := 0
	for v := range maths.PrimeFactorisation(x) {
		largestPrimeFactor = v.Value
	}
	return largestPrimeFactor
}
