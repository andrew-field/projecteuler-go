package main

import "fmt"

func main() {

	// Total sum.
	total := 0

	// Include multiple of 3.
	for i := 3; i < 1000; i += 3 {
		total += i
	}

	// Include multiples of 5.
	i := 1
	for j := 5; j < 1000; j += 5 {
		// Exclude multiples of 15.
		if i == 3 {
			i = 1
			continue
		}
		total += j
		i++
	}

	fmt.Println("Total:", total)
}
