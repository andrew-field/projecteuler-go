package main

import (
	"fmt"
	"math/big"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {

	x := big.NewInt(2)
	y := big.NewInt(1000)
	var answer big.Int
	// Calculate answer.
	answer.Exp(x, y, nil)

	digitsChannel := numbertheory.GetDigitsOfABigNumber(answer)

	total := 0

	for val := range digitsChannel {
		total += val
	}

	fmt.Println("Total:", total)
}