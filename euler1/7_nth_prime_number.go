package euler1

import (
	"context"

	"github.com/andrew-field/maths/v2"
)

// positionNthPrime returns the |n|th prime number, for |n| <= 100000
// positionNthPrime(n) returns 0 for n = 0 or n > 100000
func positionNthPrime(n int) int {
	n, err := maths.Abs(n)
	if err != nil {
		panic(err)
	}
	if n == 0 || n > 100000 {
		return 0
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	primeCh := maths.GetPrimeNumbersBelowAndIncluding(ctx, 1300000)
	for range n - 1 { // Discard the first n-1 primes.
		<-primeCh
	}

	return <-primeCh
}

// Uses brute force so as to complete in an acceptable time.
