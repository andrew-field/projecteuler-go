package main

import "fmt"

func main() {
	length := 1999999
	num := 512

	// Primes.
	primes := make([]int, length)
	for ind := range primes {
		primes[ind] = ind + 2
	}

	// Eulcidean seive.
	for ind, val := range primes {
		if val != 1 {
			for j := ind + val; j < length; j += val {
				if primes[j] != 1 {
					primes[j] = 1
				}
			}
		}
	}

	factors := make([]int, 0)

	// Make factors.
	for i := 0; i < length; i++ {
		if primes[i] != 1 && num%primes[i] == 0 {
			factors = append(factors, primes[i])
			num /= primes[i]
			i--
			if num == 1 {
				break
			}
		}
	}

	fmt.Println(factors)
	/*
		tri := 0
		for n := 1; check < 501; n++ {
			tri += n

		}
	*/
}
