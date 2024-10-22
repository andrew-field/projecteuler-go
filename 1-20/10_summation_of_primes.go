package euler1

import "github.com/andrew-field/maths"

// summationOfPrimes returns the sum of all the primes below |n|.
func summationOfPrimes(n int) int {
	sum := 0
	for val := range maths.GetPrimeNumbersBelow(n) {
		sum += val
	}

	return sum
}
