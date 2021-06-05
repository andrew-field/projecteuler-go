package euler1

import "github.com/andrew-field/maths"

// SummationOfPrimes returns the sum of all the primes below |n|.
func SummationOfPrimes(n int) int {
	sum := 0
	primeCh := maths.GetPrimeNumbersBelow(n)
	for val := range primeCh {
		sum += val
	}

	return sum
}
