/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a² + b² = c²

For example, 3² + 4² = 9 + 16 = 25 = 5².

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package euler1

// specialPythagoreanTriplet returns the product abc where a, b and c form the only Pythagorean triplet, a²+b²=c², for which a + b + c = 1000.
func specialPythagoreanTriplet() int {
	// Cycle through possible a values.
	for a := 1; ; a++ {
		num := 1000 * (500 - a)
		den := 1000 - a
		// This must be true with the two simultaneous equations.
		if num%den == 0 {
			// As there is only one triplet for which this works I don't need to make further checks on c.
			return a * num / den * (1000 - num/den - a)
		}
	}
}

// From the two initial equations you can deduce that 1000(b + a) - ba = 500,000 and b = (1000(500-a))/(1000-a). Cycle through a to find a match for a, b and then c.
