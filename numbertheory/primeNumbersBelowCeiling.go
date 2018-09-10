package numbertheory

// GetAllPrimeNumbersBelowCeiling returns all the prime numbers below and including 'ceiling'.
// Since this method uses a sieve, the max size of which is the max size of int, the ceiling
// for the largest primes to be found is also limited by the max size of int. The implementation
// conveniently handles all possible int numbers.
// This process will buffer through the channel passed in.
func GetAllPrimeNumbersBelowCeiling(primeChannel chan int, ceiling int) {

	if ceiling <= 1 {
		panic("The ceiling must be larger than 1 for any prime numbers to exist.")
	}

	// Make slice ready for primes, starting from 2.
	numbersBelowCeiling := make([]int, ceiling-1)
	for ind := range numbersBelowCeiling {
		numbersBelowCeiling[ind] = ind + 2
	}

	// Euclidean sieve.
	for ind, val := range numbersBelowCeiling {
		if val != 1 {
			primeChannel <- val
			for j := ind + val; j < ceiling-1; j += val {
				if numbersBelowCeiling[j] != 1 {
					numbersBelowCeiling[j] = 1
				}
			}
		}
	}

	close(primeChannel)
}
