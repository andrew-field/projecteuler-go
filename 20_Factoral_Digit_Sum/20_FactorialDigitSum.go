package main

import (
	"fmt"
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

func main() {

	digitsChannel := numbertheory.GetDigitsOfABigInt(*fact(big.NewInt(100)))

	total := 0

	for val := range digitsChannel {
		total += val
	}

	fmt.Println("Total:", total)
}
