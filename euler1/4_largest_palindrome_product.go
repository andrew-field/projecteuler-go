package euler1

import (
	"strconv"
)

// largestPalindromeProduct returns the largest palindrome made from the product of two 3-digit numbers.
func largestPalindromeProduct() int {
	largest := 0

	for a := 999; a > 100; a-- { // Go down from 999 to 100 instead of the other way around to generally lower the number of assignments to the largest variable.
		for b := a; b > 100; b-- {
			product := a * b
			digits := strconv.Itoa(product) // Can use a string to conveniently access the digits of the product, as it is not necessary to do any arithmetic with them.
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

// Possible to use maths.GetDigits and use slices comparison functions but in the end it would still use a loop under the hood but also provide more overhead.
