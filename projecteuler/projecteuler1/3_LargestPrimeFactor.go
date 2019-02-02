package projecteuler1

import (
	"github.com/andrew-field/testing_go/numbertheory"
)

// LargestPrimeFactor returns the largest prime factor of 600851475143.
func LargestPrimeFactor() int {
	return int(numbertheory.LargestPrimeFactor(600851475143))
}
