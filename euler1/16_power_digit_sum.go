package euler1

import (
	"math/big"

	"github.com/andrew-field/maths/v2"
)

// powerDigitSum returns the sum of the digits of the number 2¹⁰⁰⁰.
func powerDigitSum() int {
	var x *big.Int
	x.SetBit(x, 1000, 1) // Since 2¹⁰⁰⁰ is so computer friendly, use the SetBit method.

	total := 0
	for _, val := range maths.GetDigits(x) {
		total += val
	}

	return total
}

// Go makes this quite easy. Can calculate the exponent directly, but using the SetBit method is more efficient.
// x := big.NewInt(2)
// // Calculate 2¹⁰⁰⁰.
// x.Exp(x, big.NewInt(1000), nil)
