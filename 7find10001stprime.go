package main

import "fmt"
import "time"

func timetaken(start time.Time) {
	t := time.Now()
	fmt.Println("Time taken:", t.Sub(start))
}

func main() {
	// Time taken
	start := time.Now()
	defer timetaken(start)

	// Primes.
	primes := make([]int, 105000)
	for ind := range primes {
		primes[ind] = ind + 2
	}

	// Eulcidean seive.
	for i := 0; i < len(primes); i++ {
		if primes[i] != 1 {
			for j := i + 1; j < len(primes); j++ {
				if primes[j] != 1 && primes[j]%primes[i] == 0 {
					primes[j] = 1
				}
			}
		}
	}

	seek := 10001
	count := 0
	index := 0
	for ind, val := range primes {
		if val != 1 {
			count++
			index = ind
			if count == seek {
				break
			}
		}
	}

	if count < seek {
		fmt.Println("List of primes too small")
		fmt.Printf("Largest prime found is number %d: %d\n", count, primes[index])
	} else {
		fmt.Printf("Prime number %d: %d\n", seek, primes[index])
	}

}
