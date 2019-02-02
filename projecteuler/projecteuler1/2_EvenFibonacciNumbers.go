package projecteuler1

// EvenFibonacciNumbers sums all the even fibonacci numbers below 4000000.
func EvenFibonacciNumbers() int {
	// Total sum.
	total := 0

	// Fibonacci numbers 2 and 3.
	num1 := 1
	num2 := 2

	// Generate Fibonacci sequence for numbers less than 4000000.
	for num2 < 4000000 {
		total += num2
		// Only add even values so calculate three terms before adding.
		for index := 0; index < 3; index++ {
			num1, num2 = num2, num1+num2
		}
	}

	return total
}
