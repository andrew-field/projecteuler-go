package main

import "fmt"

func main() {

	squareSumTakeSumSquare := 0
	sumPart := 0

	for i := 1; i < 100; i++ {
		sumPart += i
		squareSumTakeSumSquare += (i + 1) * sumPart
	}

	fmt.Println("Difference:", 2*squareSumTakeSumSquare)
}
