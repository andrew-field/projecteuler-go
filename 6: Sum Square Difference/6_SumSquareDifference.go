package main

import "fmt"

func main() {

	// Square of sum of numbers.
	squareSum := 0

	// Sum of square of numbers.
	sumSquare := 0

	for i := 1; i < 101; i++ {
		squareSum += i
		sumSquare += i * i
	}

	// Square of sum.
	squareSum *= squareSum

	fmt.Println("Difference:", squareSum-sumSquare)
}
