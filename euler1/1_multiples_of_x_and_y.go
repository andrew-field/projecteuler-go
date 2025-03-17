package euler1

import "github.com/andrew-field/maths/v2"

// multiplesOf sums the (positive) multiples of |x| or |y| below |z|.
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

	// The number of multiples of |x| below |z|. The variable type int means that the result is floored.
	var divX int = (z - 1) / x

	// The sum of the multiples of |x| below |z|, which is equal to 1*x + 2*x + ... + divX*x => x*(1+2+...+divX). The sum of integers from 1 to n is n(n+1)/2.
	total := x * divX * (divX + 1) / 2

	// The number of multiples of |y| below |z|. The variable type int means that the result is floored.
	var divY int = (z - 1) / y

	// Add the sum of the multiples of |y| below |z| to the total so far.
	total = total + (y * divY * (divY + 1) / 2)

	lcm, err := maths.LCM(x, y)
	if err != nil {
		panic(err)
	}

	// The number of multiples of |lcm| (x and y) below |z|. The variable type int means that the result is floored.
	var divLCM int = (z - 1) / lcm

	// The sum of the multiples of |lcm| below |z|.
	sum := lcm * divLCM * (divLCM + 1) / 2

	// Subtract from the total so far, so as to not double count multiples of both x and y.
	return total - sum
}

// multiplesOf sums the (positive) multiples of |x| or |y| below |z| in the simplest way.
func multiplesOf2(x, y, z int) int {
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
	for i := min(x, y); i < z; i++ {
		if i%x == 0 || i%y == 0 {
			total += i
		}
	}

	return total
}
