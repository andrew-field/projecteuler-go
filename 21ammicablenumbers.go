package main

func power(x, y int) int {
	for num := x; y != 1; y-- {
		x *= num
	}
	return x
}

func main() {
	// Slice for primes.
	primes := make([]int, 4998)

	// Slice for numbers.
	nums := make([]int, 10001)

	// Make slice ready for prime numbers.
	for ind := range primes {
		primes[ind] = ind + 2
	}

	// Generate prime numbers through a euclidean sieve.
	for ind, val := range primes {
		if val != 1 {
			for j := ind + val; j < len(primes); j += val {
				if primes[j] != 1 {
					primes[j] = 1
				}
			}

		}
	}

	// Finish the prime slice (Get rid of the ones).
	for i := 0; primes[i] != 4999; i++ {
		if primes[i] == 1 {
			primes = append(primes[:i], primes[i+1:]...)
			i--
		}
	}

	// Go through the numbers, find all the prime factors of each number, calculate all the factors of the number,
	// store the sum of all the factors at the location in the number array. ind 1 == number 1, the value stored there
	// is d(1).
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
		if num != 1 {
			factors = append(factors, 0)
		}
		// total := 0

	}
}
