package projecteuler1

// SumSquareDifference returns the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
func SumSquareDifference() int {
	squareSumTakeSumSquare := 0
	sumPart := 0

	// Some rearranging can make the calculation into this simplest form.
	for i := 1; i < 100; i++ {
		sumPart += i
		squareSumTakeSumSquare += (i + 1) * sumPart
	}

	return 2 * squareSumTakeSumSquare
}
