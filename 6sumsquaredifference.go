package main

import "fmt"

func main() {
	// Numbers 1 to 100.
	num := make([]int, 100)

	// Sum of numbers.
	total := 0

	// Sum of squares.
	sumsquare := 0

	for ind := range num {
		num[ind] = ind + 1
		total += num[ind]
		sumsquare += num[ind] * num[ind]
	}

	// Square of sum.
	squaresum := total * total

	fmt.Println("Difference:", squaresum-sumsquare)

}
