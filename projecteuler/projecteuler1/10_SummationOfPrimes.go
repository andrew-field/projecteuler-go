package projecteuler1

import (
	"github.com/andrew-field/testing_go/numbertheory"
)

// SummationOfPrimes returns the sum of all the primes below two million.
func SummationOfPrimes() int {
	primeChannel := numbertheory.GetAllPrimeNumbersBelowCeiling(2000000)

	sum := 0
	for val := range primeChannel {
		sum += val
	}

	return sum
}
