package euler1

import "github.com/andrew-field/maths"

// PositionNthPrime returns the |n|th prime number, for |n| <= 100000
// PositionNthPrime(x) returns 0 for x = 0, x > 100000
func PositionNthPrime(n int) int {
	n = maths.Abs(n)
	if n == 0 || n > 100000 {
		return 0
	}

	primeCh := maths.GetPrimeNumbersBelow(1300000)
	for i := 1; i < n; i++ {
		<-primeCh
	}

	return <-primeCh
}

// Uses brute force so as to complete in an acceptable time.
