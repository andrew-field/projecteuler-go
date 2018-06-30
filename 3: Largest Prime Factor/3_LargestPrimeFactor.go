package main

import "fmt"
import "time"

func main() {

	// Number in question.
	num := 600851475143

	// Upper limit for size of prime numbers.
	length := 9999
	primes := make([]int, length)

	// Make slice ready for primes, starting from 2.
	for ind := range primes {
		primes[ind] = ind + 2
	}

	factors := make([]int, 0)

	test := time.Now()
	// Make prime splice.
	// Euclidean sieve.
	for ind, val := range primes {
		if val != 1 {
			// For each prime, test to see if it is a factor and for how many times.
			for num%val == 0 {
				factors = append(factors, val)
				num /= val
			}
			// If found all factors, break.
			if num == 1 {
				break
			}
			// Generate primes.
			for j := ind + val; j < length; j += val {
				if primes[j] != 1 {
					primes[j] = 1
				}
			}
		}
	}
	fmt.Println("Sieve time:", time.Since(test))

	fmt.Println("Largest:", factors[len(factors)-1])

}
