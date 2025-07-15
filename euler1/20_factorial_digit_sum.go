/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

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

	z := new(big.Int)
	z.MulRange(2, int64(absN))

	total := 0
	for _, val := range maths.GetDigits(z) {
		total += val
	}

	return total
}

// The standard library maths package does the heavy lifting.
