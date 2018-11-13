package main

import (
	"fmt"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {

	// Keep track of largest.
	largest := 0

	for a := 100; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			product := a * b
			digits := numbertheory.GetDigitsOfANumberInSlice(product)
			if lastIndex := len(digits) - 1; digits[0] == digits[lastIndex] && digits[1] == digits[lastIndex-1] && digits[2] == digits[lastIndex-2] {
				if product > largest {
					largest = product
				}
			}
		}
	}

	fmt.Println("Largest palindrome:", largest)
}
