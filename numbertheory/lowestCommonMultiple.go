package numbertheory

import "sync"

// LowestCommonMultiple returns the lowest common multiple (lcm) of a group of numbers.
func LowestCommonMultiple(numbers ...uint) uint {

	// Slice of slices containing the prime factorisation of each number.
	primeFactorisations := make([][]uint, 0)

	// Fill slices
	for _, val := range numbers {
		newFactorisation := make([]uint, 0)
		factorChannel := make(chan uint, 100)
		// Should not need syncing as the closing of the primeFactorChannel is the very last operation
		// of GetPrimeFactorisation.

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			go GetPrimeFactorisation(factorChannel, val)
		}()
		for val := range factorChannel {
			newFactorisation = append(newFactorisation, val)
		}
		primeFactorisations = append(primeFactorisations, newFactorisation)
		wg.Wait()
	}

	// For each prime number, any duplicates are removed from the other factorisations (once).
	// This is so each prime number appears the correct number of times for the lcm calculation.
	// Another way to think about it is that each prime number appears exactly the maximum number
	// of times it appears in any individual factorisation.
	// See lcm calculation.
	for ind, val := range primeFactorisations[:len(primeFactorisations)-1] {
		for _, prime := range val {
			for ind2, val2 := range primeFactorisations[ind+1:] {
				for ind3, val3 := range val2 {
					if val3 == prime {
						primeFactorisations[ind+1+ind2][len(val2)-1], primeFactorisations[ind+1+ind2][ind3] = primeFactorisations[ind+1+ind2][ind3], primeFactorisations[ind+1+ind2][len(val2)-1]
						primeFactorisations[ind+1+ind2] = primeFactorisations[ind+1+ind2][:len(val2)-1]
						break
					}
				}
			}
		}
	}

	// Multiply the remaining primes together.
	var product uint = 1
	for _, val := range primeFactorisations {
		for _, prime := range val {
			product *= prime
		}
	}

	return product
}
