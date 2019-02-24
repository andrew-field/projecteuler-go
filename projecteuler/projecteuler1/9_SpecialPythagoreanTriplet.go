package projecteuler1

// SpecialPythagoreanTriplet returns the product abc where a, b and c form the only Pythagorean triplet for which a + b + c = 1000.
func SpecialPythagoreanTriplet() int {
	// Cycle through possible a's and b's.
	for a := 1; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			// This must be true with the two simultaneous equations.
			if 500000 == 1000*(b+a)-b*a {
				// As there is only one triplet for which this works I don't need to make further checks on c.
				return a * b * (1000 - b - a)
			}
		}
	}

	return 0
}
