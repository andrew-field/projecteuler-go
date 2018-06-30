package main

import (
	"fmt"
	"math/big"
)

// Factorial function.
func fact(n *big.Int) *big.Int {
	if !n.IsInt64() || n.Int64() != 2 {
		c := big.NewInt(1)
		return n.Mul(n, fact(c.Sub(n, c)))
	}
	return big.NewInt(2)
}

func main() {

	// Fact 100.
	a := fact(big.NewInt(100))

	ten := big.NewInt(10)
	digit := big.NewInt(0)

	var total int64

	// Dividing these Ints by 10 truncates the decimal places so it is fine. Reverse order of digits; doesn't matter.
	for !a.IsInt64() || a.Int64() > 9 {
		a.DivMod(a, ten, digit)
		total += digit.Int64()
	}
	total += a.Int64()
	fmt.Println("Answer:", total)
}
