package euler1

import "github.com/andrew-field/maths/v2"

// multiplesOf sums the (positive) multiples of |x| and |y| below |z| in the simplest way.
func multiplesOf(x, y, z int) int {
	x, err := maths.Abs(x)
	if err != nil {
		panic(err)
	}
	y, err = maths.Abs(y)
	if err != nil {
		panic(err)
	}
	z, err = maths.Abs(z)
	if err != nil {
		panic(err)
	}

	total := 0
	for i := 1; i < z; i++ {
		if i%x == 0 || i%y == 0 {
			total += i
		}
	}

	return total
}
