package euler1

import (
	"math/big"

	"github.com/andrew-field/maths"
)

// PowerDigitSum returns the sum of the digits of the number 2¹⁰⁰⁰.
func PowerDigitSum() int {

	x := big.NewInt(2)
	// Calculate 2¹⁰⁰⁰.
	x.Exp(x, big.NewInt(1000), nil)

	total := 0
	for val := range maths.DigitsBig(x) {
		total += val
	}

	return total
}

// Go makes this quite easy. I don't know if there is a clever way of doing this. Maybe using bits.
