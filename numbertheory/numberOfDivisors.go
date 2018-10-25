package numbertheory

// GetNumberOfDivisors returns the number of divisors of number.
func GetNumberOfDivisors(number uint) int {

	if number <= 1 {
		panic("The number must be larger than 1")
	}

	primeFactorChannel := GetPrimeFactorisation(number)

	// Calculate the number of divisors.
	power := 0
	var check uint
	divisors := 1
	for val := range primeFactorChannel {
		if val == check {
			power++
		} else {
			divisors *= power + 1 // The first time this is OK because power = 0.
			power = 1
			check = val
		}
	}

	divisors *= power + 1

	return divisors
}
