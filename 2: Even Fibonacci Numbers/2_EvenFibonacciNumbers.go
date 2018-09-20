package main

import "fmt"

func main() {

	// Fibonacci numbers 2 and 3.
	num1 := 1
	num2 := 2

	// Sum.
	total := 0

	// Generate Fibonacci.
	for num2 < 4000000 {
		total += num2
		num1, num2 = num2, num1+num2
		num1, num2 = num2, num1+num2
		num1, num2 = num2, num1+num2
	}

	fmt.Println("Total:", total)
}
