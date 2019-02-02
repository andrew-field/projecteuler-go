package numbertheory

// GetNthPrimeNumber returns the prime number at the nth position. E.g. the 12th prime number is 37.
func GetNthPrimeNumber(position uint) uint {
	if position == 0 {
		panic("The position must be larger than 0.")
	}

	// Make prime and done channel.
	primeChannel, doneChannel := GetPrimeNumbersContinuously(100)

	var count uint = 1
	var prime uint
	for ; count <= position; count++ {
		prime = <-primeChannel
	}
	doneChannel <- true

	return prime
}
