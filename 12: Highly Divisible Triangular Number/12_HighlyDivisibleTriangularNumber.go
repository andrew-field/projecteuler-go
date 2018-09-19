package main

import (
	"fmt"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {

	// Each triangular number.
	var triangularNumber uint = 1

	// Calculate the number of divisors and see if it is greater than 500.
	var i uint = 2
	for divisors := 1; divisors < 501; i++ {
		triangularNumber += i

		divisors = numbertheory.GetNumberOfDivisors(triangularNumber)
	}

	fmt.Println("Triangular number: ", triangularNumber)
}
