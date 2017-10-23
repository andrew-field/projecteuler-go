package main

import "fmt"

func main() {

	// Upper limit for size of prime numbers.
	length := 199999

	// Make slice for primes starting from 2.
	primes := make([]int, length)
	for ind := range primes {
		primes[ind] = ind + 2
	}

	// Euclidean sieve.
	for ind, val := range primes {
		if val != 1 {
			for j := ind + val; j < length; j += val {
				if primes[j] != 1 {
					primes[j] = 1
				}
			}
		}
	}

	// Each triangular number.
	tri := 1

	// Number of divisors.
	divisors := 1

	for n := 2; divisors < 501; n++ {
		tri += n

		temp := tri
		factors := make([]int, 0)

		// Make factors.
		for _, val := range primes {
			if val != 1 {
				for temp%val == 0 {
					factors = append(factors, val)
					temp /= val
				}
				if temp == 1 {
					break
				}
			}
		}
		// If the above for did not break, then there were not enough prime numbers to find all the factors for the current triangular number.
		if temp != 1 {
			fmt.Println("Not enough primes!")
			break
		}

		// Calculate the number of divisors.
		power := 0
		check := factors[0]
		divisors = 1

		for _, val := range factors {
			if val == check {
				power++
			} else {
				divisors *= power + 1
				power = 1
				check = val
			}
		}
		divisors *= power + 1
	}

	fmt.Println("Tri: ", tri)
}
