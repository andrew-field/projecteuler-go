package euler1

import "math/big"

// latticePaths returns the number of routes through a 20x20 grid starting in the top left, finishing in the bottom right
// and only moving right or down.
func latticePaths() int {
	var z big.Int
	return int(z.Binomial(40, 20).Int64())
}

// This problem is just picking k choices from n scenarios, where order does not matter.
// Hence, use the binomial coefficient. The problem can be thought of by making 40 steps/paths with only 20 real independent choices.
// This program is just used to calculate the binomial coefficient.

// I then found: "The number of lattice paths from the origin (0,0) to a point (a,b) is the binomial coefficient (a+b; a) (Hilton and Pedersen 1991)."
