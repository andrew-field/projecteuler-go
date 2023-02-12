package euler1

import "github.com/andrew-field/maths"

// SummationOfPrimes returns the sum of all the primes below |n|.
func SummationOfPrimes(n int) int {
	sum := 0
	for val := range maths.GetPrimeNumbersBelow(n) {
		sum += val
	}

	 return sum
}
