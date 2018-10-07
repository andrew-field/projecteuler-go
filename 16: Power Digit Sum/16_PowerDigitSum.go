package main

import "math/big"
import "fmt"

func main() {

	x := big.NewInt(2)
	y := big.NewInt(1000)
	var z big.Int
	// Calculate answer.
	z.Exp(x, y, nil)

	// Reuse y = 10.
	y.SetInt64(10)

	var total int64

	// Dividing these Ints by 10 truncates the decimal places so it is fine.
	// Reverse order of digits; doesn't matter. Reuse x.
	for z.Int64() > 9 || !z.IsInt64() {
		// Go is so nice with this handy function that does everything I need.
		// Sets z to z / y (10) and sets x to z mod y (The last digit).
		z.DivMod(&z, y, x)
		total += x.Int64()
	}

	total += z.Int64()
	fmt.Println("Answer:", total)
}
