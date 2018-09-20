package numbertheory

// GetNumberOfDivisors returns the number of divisors of number.
func GetNumberOfDivisors(number uint) int {
	primeFactorChannel := make(chan uint, 100)

	// Should not need syncing as the closing of the primeFactorChannel is the very last operation
	// of GetPrimeFactorisation.
	go GetPrimeFactorisation(primeFactorChannel, number)

	// Calculate the number of divisors.
	power := 0
	var check uint
	divisors := 1
	for val := range primeFactorChannel {
		if val == check {
			power++
		} else {
			divisors *= power + 1
			power = 1
			check = val
		}
	}

	divisors *= power + 1

	return divisors
}
