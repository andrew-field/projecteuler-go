package main

import "fmt"

func main() {

	num := 600851475143

	length := 9999
	primes := make([]int, length)

	for ind := range primes {
		primes[ind] = ind + 2
	}

	// Make prime splice.
	// Euclidean seive.
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

	largest := factors[0]

	for _, val := range factors {
		if val > largest {
			largest = val
		}
	}

	fmt.Println(largest)

}
