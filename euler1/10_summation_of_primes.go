/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package euler1

import (
	"context"

	"github.com/andrew-field/maths/v2"
)

// summationOfPrimes returns the sum of all the primes below and including |n|.
func summationOfPrimes(n int) int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sum := 0
	for val := range maths.GetPrimeNumbersBelowAndIncluding(ctx, n) {
		sum += val
	}

	return sum
}
