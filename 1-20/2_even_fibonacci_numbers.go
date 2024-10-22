package euler1

// evenFibonacciNumbers sums all the even fibonacci numbers below x.
func evenFibonacciNumbers(x int) int {
	total := 0

	// Fibonacci numbers 2 and 3.
	num1, num2 := 1, 2

	// Generate Fibonacci sequence for numbers less than x.
	for num2 < x {
		total += num2
		// Only add even values so calculate three terms before adding.
		for i := 0; i < 3; i++ {
			num1, num2 = num2, num1+num2
		}
	}

	return total
}

// The Fibonacci numbers are on a cycle explained below.

//  1. odd + even = odd. --> 2.
//  2. even + odd = odd. --> 3.
//  3. odd + odd = even. --> back to 1.

// Hence you only need to sum up every third Fibonacci number for the even numbers.
