/*
2¹⁵ = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2¹⁰⁰⁰?
*/

package euler1

import (
	"math/big"

	"github.com/andrew-field/maths/v2"
)

// powerDigitSum returns the sum of the digits of the number 2¹⁰⁰⁰.
func powerDigitSum() int {
	x := new(big.Int)
	x.SetBit(x, 1000, 1) // Since 2¹⁰⁰⁰ is so computer friendly, use the SetBit method.

	total := 0
	for _, val := range maths.GetDigits(x) {
		total += val
	}

	return total
}

// Go makes this quite easy. Can calculate the exponent directly using Exp(x, y, m *Int) but using the SetBit method is more efficient.
// I don't know a clever way of doing this.
