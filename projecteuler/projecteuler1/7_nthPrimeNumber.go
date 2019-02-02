package projecteuler1

import (
	"github.com/andrew-field/testing_go/numbertheory"
)

// Position10001Prime returns the 10001st prime number.
func Position10001Prime() int {
	return int(numbertheory.GetNthPrimeNumber(10001))
}
