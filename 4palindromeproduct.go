package main

import "fmt"

// Assume it is 6 digits and the two numbers are greater than 400.
// 456/10 = 45 with int.

func main() {
	largest := 0

	for a := 400; a < 999; a++ {
		for b := 400; b < 999; b++ {
			n := a * b
			if n%10 == n/100000 && (n/10)%10 == (n/10000)%10 && (n/100)%10 == (n/1000)%10 {
				if n > largest {
					largest = n
				}
			}
		}
	}

	fmt.Println("Largest:", largest)
}
