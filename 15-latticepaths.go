package main

import "fmt"
import "math/big"

// This program is just used to calculate the binomial coefficient. Thsé problem is just picking k choices from n scenarios, where order does not matter.
// Hence, binomal coefficient. The problem can be thought of by making 40 steps/paths with only 20 real independent choices.

func main() {
	fmt.Println("Answer:", big.NewInt(0).Binomial(40, 20).Int64())
}
