package main

import "fmt"

func main() {
	// Slice for primes.
	primes := make([]int, 96)

	// Slice for numbers.
	nums := make([]int, 10001)

	// Make slice ready for prime numbers.
	for ind := range primes {
		primes[ind] = ind + 2
	}

	// Generate prime numbers through a euclidean sieve.
	for ind, val := range primes {
		if val != 1 {
			for j := ind + val; j < 96; j += val {
				if primes[j] != 1 {
					primes[j] = 1
				}
			}

		}
	}

	// Finish the prime array slices (Get rid of the ones).
	for i := 0; primes[i] != 97; i++ {
		if primes[i] == 1 {
			primes = append(primes[:i], primes[i+1:]...)
			i--
		}
	}

	for ind := range nums {
		num := ind + 1
		factors := make([]int, 0)
		for _, val := range primes {
			for num%val == 0 {
				factors = append(factors, val)
				num /= val
			}
			if num == 1 {
				break
			}
		}
		total := 0

	}

	fmt.Println(primes)
}
