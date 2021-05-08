package projecteuler1

import (
	"math/big"

	"github.com/andrew-field/testing_go/numbertheory"
)

// PowerDigitSum returns the sum of the digits of the number 2**1000.
func PowerDigitSum() int {

	x := big.NewInt(2)
	y := big.NewInt(1000)
	// Calculate answer.
	x.Exp(x, y, nil)

	digitsChannel := numbertheory.GetDigitsOfABigInt(x)
	total := 0
	for val := range digitsChannel {
		total += val
	}

	return total
}
