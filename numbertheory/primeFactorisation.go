package numbertheory

import "sync"

// GetPrimeFactorisation fills a channel with the prime factors of a number.
func GetPrimeFactorisation(primeFactorChannel chan uint, numberToFactorise uint) {

	if numberToFactorise <= 1 {
		panic("The number to factorise must be larger than 1")
	}

	// Make prime and done channel.
	primeChannel := make(chan uint, 100)
	doneChannel := make(chan bool, 1)

	// Generate primes.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		GetPrimeNumbersContinuously(primeChannel, doneChannel, 100)
	}()

	// For each prime, see if it is a factor.
	for val := range primeChannel {
		for numberToFactorise%val == 0 {
			primeFactorChannel <- val
			numberToFactorise /= val
		}

		// If found all factors, break.
		if numberToFactorise == 1 {
			doneChannel <- true
			close(doneChannel)
			break
		}
	}

	// Check
	if numberToFactorise != 1 {
		panic("Could not factorise.")
	}

	// Flush prime channel to prevent blocking.
	for len(primeChannel) > 0 {
		<-primeChannel
	}

	wg.Wait()

	close(primeFactorChannel)
}
