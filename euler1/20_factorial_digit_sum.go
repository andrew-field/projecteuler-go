package euler1

import (
	"math/big"

	"github.com/andrew-field/maths/v2"
)

// factorialDigitSum returns the sum of the digits in the number |n|!
func factorialDigitSum(n int) int {
	absN, err := maths.Abs(n)
	if err != nil {
		panic(err)
	}

	var z *big.Int
	z.MulRange(2, int64(absN))

	total := 0
	for _, val := range maths.GetDigits(z) {
		total += val
	}

	return total
}

// The standard library maths package does the heavy lifting.
