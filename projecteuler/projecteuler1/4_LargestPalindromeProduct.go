package projecteuler1

import (
	"github.com/andrew-field/testing_go/numbertheory"
)

// LargestPalindromeProduct returns the largest palindrome made from the product of two 3-digit numbers.
func LargestPalindromeProduct() int {
	// Keep track of largest.
	largest := 0

	for a := 100; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			product := a * b
			digits := numbertheory.GetDigitsOfAnIntInSlice(product)
			lastIndex := len(digits) - 1
			// General palindrome checker.
			for index := 0; digits[index] == digits[lastIndex-index]; index++ {
				// Works fine with even and odd length as int/2 is always the lower bound.
				if index == lastIndex/2 {
					if product > largest {
						largest = product
					}
					break
				}
			}
		}
	}

	return largest
}
