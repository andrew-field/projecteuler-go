package main

import "fmt"
import "time"

func main() {

	// Upper limit for size of prime numbers.
	length := 104999

	// Make slice ready for primes starting from 2.
	primes := make([]int, length)
	for ind := range primes {
		primes[ind] = ind + 2
	}

	count := 0

	start := time.Now()
	// Euclidean sieve.
	for ind, val := range primes {
		if val != 1 {
			count++
			if count == 10001 {
				fmt.Printf("Prime number %d: %d\n", count, val)
				break
			} else {
				for j := ind + val; j < length; j += val {
					if primes[j] != 1 {
						primes[j] = 1
					}
				}
			}
		}
	}
	fmt.Println("Sieve time:", time.Since(start))
}
