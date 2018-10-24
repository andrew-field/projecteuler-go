package main

import (
	"fmt"
	"math/big"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {
	answer := big.NewInt(123)

	digitsChannel := numbertheory.GetDigitsOfABigNumber(*answer)

	total := 0

	for val := range digitsChannel {
		fmt.Println(val)
		total += val
	}

	fmt.Println("Total:", total)
}
