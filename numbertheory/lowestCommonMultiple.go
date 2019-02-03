package numbertheory

// LowestCommonMultiple returns the lowest common multiple (LCM) of a group of numbers.
func LowestCommonMultiple(allNumbers ...uint) uint {

	// Slice of slices containing the prime factorisation of each number.
	primeFactorisations := make([][]uint, 0)

	// Fill slices
	for _, val := range allNumbers {
		// Special case if a number is 0.
		if val == 0 {
			return 0
		}

		// Special case if a number is 1. This function could handle ones as it is but this is more efficient.
		if val == 1 {
			continue
		}

		primeFactorisations = append(primeFactorisations, GetPrimeFactorisationInSlice(val))
	}

	// Special case if all the numbers are 1.
	if len(primeFactorisations) == 0 {
		return 1
	}

	// For each prime number, any duplicates are removed from the other factorisations (once).
	// This is so each prime number appears the correct number of times for the LCM calculation.
	// Another way to think about it is that each prime number appears exactly the maximum number
	// of times it appears in any of the individual factorisations.
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
