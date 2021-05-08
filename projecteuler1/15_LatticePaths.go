package projecteuler1

import "math/big"

// This problem is just picking k choices from n scenarios, where order does not matter.
// Hence, use the binomal coefficient. The problem can be thought of by making 40 steps/paths with only 20 real independent choices.
// This program is just used to calculate the binomial coefficient.

// LatticePaths returns the number of routes through a 20x20 grid starting in the top left, finishing in the bottom right
// and only moving right or down.
func LatticePaths() int {
	z := big.NewInt(0)
	return int(z.Binomial(40, 20).Int64())
}
