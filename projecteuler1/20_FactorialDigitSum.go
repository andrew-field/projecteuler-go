package projecteuler1

import (
	"math/big"

	"github.com/andrew-field/testing_go/numbertheory"
)

// Factorial function.
func fact(number *big.Int) *big.Int {
	if !number.IsInt64() || number.Int64() != 2 {
		one := big.NewInt(1)
		return number.Mul(number, fact(one.Sub(number, one)))
	}
	return big.NewInt(2)
}

// FactorialDigitSum returns the sum of the digits in the number 100!
func FactorialDigitSum() int {
	digitsChannel := numbertheory.GetDigitsOfABigInt(fact(big.NewInt(100)))

	total := 0
	for val := range digitsChannel {
		total += val
	}

	return total
}
