package main

import "fmt"
import "time"

func main() {
	// Time taken
	start := time.Now()

	length := 1999998

	// Primes.
	primes := make([]int, length)
	for ind := range primes {
		primes[ind] = ind + 2
	}

	total := 0
	// Eulcidean seive.
	for ind, val := range primes {
		if val != 1 {
			total += val
			for j := ind + val; j < length; j += val {
				if primes[j] != 1 {
					primes[j] = 1
				}
			}
		}
	}

	fmt.Println("Total", total)
	fmt.Println("Time taken:", time.Now().Sub(start))
}
