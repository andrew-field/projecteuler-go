package numbertheory

// GetPrimeFactorisation fills a channel with the prime factors of a number.
// Syncing and safely exiting this function should be done through flushing the prime factor channel.
func GetPrimeFactorisation(primeFactorChannel chan uint, numberToFactorise uint) {

	if numberToFactorise <= 1 {
		panic("The number to factorise must be larger than 1")
	}

	// Make prime and done channel.
	primeChannel := make(chan uint, 100)
	doneChannel := make(chan bool)

	go GetPrimeNumbersContinuously(primeChannel, doneChannel, 100)

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
	close(doneChannel)

	close(primeFactorChannel)
}
