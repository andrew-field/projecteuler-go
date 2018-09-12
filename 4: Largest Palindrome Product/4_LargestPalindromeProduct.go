package main

import (
	"fmt"
	"math"
)

// Assume it is 6 digits and the two numbers are greater than 400.
// 456/10 = 45 with int.

func main() {

	// Keep track of largest.
	largest := 0

	for a := 100; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			n := a * b
			numberOfDigitsMinusOne := math.Floor(math.Log10(float64(n)))
			maxExponentOf10 := int(math.Pow10(int(numberOfDigitsMinusOne)))

			if n%10 == n/maxExponentOf10 && (n/10)%10 == (n/(maxExponentOf10/10))%10 && (n/100)%10 == (n/(maxExponentOf10/100))%10 {
				if n > largest {
					largest = n
				}
			}
		}
	}

	fmt.Println("Largest palindrome:", largest)
}
