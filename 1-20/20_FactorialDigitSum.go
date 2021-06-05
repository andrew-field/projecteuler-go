package euler1

import (
	"math/big"

	"github.com/andrew-field/maths"
)

// FactorialDigitSum returns the sum of the digits in the number |n|!
func FactorialDigitSum(n int) int {

	var z big.Int
	total := 0
	for val := range maths.DigitsBig(z.MulRange(2, int64(maths.Abs(n)))) {
		total += val
	}

	return total
}
