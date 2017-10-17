package main

import "fmt"
import "time"

func main() {

	length := 1999998

	// Primes.
	primes := make([]int, length)
	for ind := range primes {
		primes[ind] = ind + 2
	}

	total := 0
	test := time.Now()
	// Euclidean sieve.
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
	fmt.Println("Sieve time:", time.Since(test))
}
