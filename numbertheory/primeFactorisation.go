package numbertheory

// GetPrimeFactorisation fills a channel with the prime factors of a number.
// Syncing and safely exiting this function can be done through flushing the prime factor channel.
func GetPrimeFactorisation(numberToFactorise uint) chan uint {
	primeFactorChannel := make(chan uint, 100)

	// Special case I will use for 0 and 1.
	if numberToFactorise == 0 || numberToFactorise == 1 {
		primeFactorChannel <- numberToFactorise
		close(primeFactorChannel)
	} else {
		go func() {
			// Make prime and done channel.
			primeChannel, doneChannel := GetPrimeNumbersContinuously(100)

			// For each prime, see if it is a factor.
			for val := range primeChannel {
				for numberToFactorise%val == 0 {
					primeFactorChannel <- val
					numberToFactorise /= val
				}

				// If found all factors then finish.
				if numberToFactorise == 1 {
					doneChannel <- true
					close(primeFactorChannel)
					return
				}
			}
		}()
	}

	return primeFactorChannel
}

// GetPrimeFactorisationInASlice returns a slice containing the prime factorisation of a number.
// This is useful when the whole factorisation is required before further calculation.
// The special case of 0 and 1 follows from "GetPrimeFactorisation" and will be 0 and 1 respectively.
func GetPrimeFactorisationInASlice(numberToFactorise uint) []uint {
	return channelIntoSlice(GetPrimeFactorisation, numberToFactorise)
}

// LargestPrimeFactor returns the largest prime factor of a number.
// The special case of 0 and 1 follows from "GetPrimeFactorisation" and will be 0 and 1 respectively.
func LargestPrimeFactor(number uint) uint {
	primeFactorChannel := GetPrimeFactorisation(number)

	var largestPrimeFactor uint
	for val := range primeFactorChannel {
		largestPrimeFactor = val
	}

	return largestPrimeFactor
}
