package numbertheory

// GetPrimeFactorisation fills a channel with the prime factors of a number.
func GetPrimeFactorisation(primeFactors chan uint, numberToFactorise uint) {

	if numberToFactorise <= 1 {
		panic("The number to factorise must be larger than 1")
	}

	// Make prime and done channel.
	primeChannel := make(chan uint, 100)
	doneChannel := make(chan bool, 1)

	// Generate primes.
	go GetPrimeNumbersContinuously(primeChannel, 1000, doneChannel)

	for val := range primeChannel {
		for numberToFactorise%val == 0 {
			primeFactors <- val
			numberToFactorise /= val
		}

		// If found all factors, break.
		if numberToFactorise == 1 {
			doneChannel <- true
			break
		}
	}

	// Check
	if numberToFactorise != 1 {
		panic("Could not factorise.")
	}

	close(primeFactors)
}
