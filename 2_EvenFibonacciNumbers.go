package main

import "fmt"

func main() {

	// Fibonacci numbers 1 and 2.
	num1 := 1
	num2 := 1

	// Sum.
	total := 0

	// Generate Fibonacci.
	for num1 < 4000000 {
		num2 += num1
		if num2%2 == 0 {
			total += num2
		}
		num1, num2 = num2, num1
	}

	fmt.Println(total)
}
