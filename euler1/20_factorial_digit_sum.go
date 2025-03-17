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

	total := 0
	var z big.Int
	for _, val := range maths.GetDigitsBig(z.MulRange(2, int64(absN))) {
		total += val
	}

	return total
}

// The go math package does the heavy lifting.
