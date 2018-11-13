package main

import "fmt"
import "math/big"

// This problem is just picking k choices from n scenarios, where order does not matter.
// Hence, use the binomal coefficient. The problem can be thought of by making 40 steps/paths with only 20 real independent choices.
// This program is just used to calculate the binomial coefficient.

func main() {
	var z big.Int
	fmt.Println("Answer:", z.Binomial(40, 20))
}
