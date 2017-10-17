package main

import "fmt"
import "time"

func main() {

	num := 600851475143

	length := 9999
	primes := make([]int, length)

	for ind := range primes {
		primes[ind] = ind + 2
	}

	factors := make([]int, 0)

	test := time.Now()
	// Make prime splice.
	// Euclidean sieve.
	for ind, val := range primes {
		if val != 1 {
			for num%val == 0 {
				factors = append(factors, val)
				num /= val
			}
			if num == 1 {
				break
			}
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
