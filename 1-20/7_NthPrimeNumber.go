package euler1

import "github.com/andrew-field/maths"

// PositionNthPrime returns the |n|th prime number.
// PositionNthPrime(0) returns 0.
func PositionNthPrime(n int) int {
	n = maths.Abs(n)
	if n == 0 {
		return 0
	}

	primeCh, doneCh := maths.GetPrimeNumbers()
	defer func() {
		doneCh <- true
	}()
	for i := 1; i < n; i++ {
		<-primeCh
	}

	return <-primeCh
}
