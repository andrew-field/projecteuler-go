package main

import "fmt"

func methodOne() {
	// Total sum.
	total := 0

	// Include multiple of 3.
	for i := 3; i < 1000; i += 3 {
		total += i
	}

	// Include multiples of 5.
	i := 1
	for j := 5; j < 1000; j += 5 {
		// Exclude multiples of 3.
		if i == 3 {
			i = 1
			continue
		}
		total += j
		i++
	}

	fmt.Println("Total:", total)
}

func methodTwo() {
	// Total sum.
	total := 0

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	fmt.Println("Total:", total)
}

func main() {
	methodOne()
}
