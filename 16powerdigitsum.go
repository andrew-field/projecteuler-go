package main

import "math/big"
import "fmt"

func main() {
	x := big.NewInt(2)
	y := big.NewInt(1000)
	z := big.NewInt(0).Exp(x, y, nil)

	digits := make([]int64, 0)
	// Reuse y.
	ten := big.NewInt(10)

	// Dividing these Ints by 10 truncates the decimal placesnso it is fine. Reverse order of digits; doesn't matter. Reuse x.
	for z.Int64() > 9 || !z.IsInt64() {
		z.DivMod(z, ten, x)
		digits = append(digits, x.Int64())
	}
	digits = append(digits, z.Int64())

	var total int64
	for _, val := range digits {
		total += val
	}

	fmt.Println("Answer:", total)
}
