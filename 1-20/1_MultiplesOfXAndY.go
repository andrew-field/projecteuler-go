package euler1

import "github.com/andrew-field/maths"

// MultiplesOf sums the (positive) multiples of |x| and |y| below |z| in the simplest way.
func MultiplesOf(x, y, z int) int {
	x, y, z = maths.Abs(x), maths.Abs(y), maths.Abs(z)

	total := 0
	for i := 1; i < z; i++ {
		if i%x == 0 || i%y == 0 {
			total += i
		}
	}

	return total
}
