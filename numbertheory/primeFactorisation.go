package numbertheory

// GetPrimeFactorisation fills a channel with the prime factors of a number.
// Syncing and safely exiting this function can be done through flushing the prime factor channel.
func GetPrimeFactorisation(numberToFactorise uint) chan uint {
	primeFactorChannel := make(chan uint, 100)

	go func() {
		// Make prime and done channel.
		primeChannel, doneChannel := GetPrimeNumbersContinuously(100)

		// For each prime, see if it is a factor.
		for val := range primeChannel {
			for numberToFactorise%val == 0 {
				primeFactorChannel <- val
				numberToFactorise /= val
			}

			// If found all factors, break.
			if numberToFactorise == 1 {
				doneChannel <- true
				break
			}
		}
		<-doneChannel

		close(primeFactorChannel)
	}()

	return primeFactorChannel
}

// LargestPrimeFactor returns the largest prime factor of a number.
func LargestPrimeFactor(number uint) uint {
	primeFactorChannel := GetPrimeFactorisation(number)

	var largestPrimeFactor uint
	for val := range primeFactorChannel {
		largestPrimeFactor = val
	}

	return largestPrimeFactor
}
