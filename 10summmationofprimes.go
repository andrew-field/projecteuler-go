package main

import "fmt"
import "time"

func main() {
	// Time taken
	start := time.Now()

	// Primes.
	primes := make([]int, 1999998)
	for ind := range primes {
		primes[ind] = ind + 2
	}

	total := 0
	// Eulcidean seive.
	for i := 0; i < len(primes); i++ {
		if primes[i] != 1 {
			total += primes[i]
			for j := i + primes[i]; j < len(primes); j += primes[i] {
				if primes[j] != 1 {
					primes[j] = 1
				}
			}
		}
	}

	fmt.Println("Total", total)
	fmt.Println("Time taken:", time.Now().Sub(start))
}
