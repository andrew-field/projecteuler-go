package euler1

import "testing"

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []struct {
		desc           string
		actualResult   int
		expectedResult int
	}{
		{"1a: Multiples of 3 and 5", multiplesOf(3, 5, 1000), 233168},
		{"1b: Multiples of 3 and 5", multiplesOf2(3, 5, 1000), 233168},
		{"2a: Even Fibonacci Numbers", evenFibonacciNumbers(4000000), 4613732},
		{"2b: Even Fibonacci Numbers", evenFibonacciNumbers2(4000000), 4613732},
		{"3: Largest Prime Factor", largestPrimeFactor(600851475143), 6857},
		{"4: Largest Palindrome Product", largestPalindromeProduct(), 906609},
		{"5: Smallest Multiple", smallestMultiple(), 232792560},
		{"6: Sum Square Difference", sumSquareDifference(1, 100), 25164150},
		{"7: 10001st Prime Number", positionNthPrime(10001), 104743},
		{"8: Largest Product in a Series", largestProductInASeries(), 23514624000},
		{"9: Special Pythagorean triplet", specialPythagoreanTriplet(), 31875000},
		{"10: Summation of Primes", summationOfPrimes(2000000), 142913828922},
		{"11: Largest Product in a Grid", largestProductInAGrid(), 70600674},
		{"12: Highly Divisible Triangular Number", highlyDivisibleTriangularNumber(500), 76576500},
		{"13: Large Sum", largeSum(), 5537376230},
		{"14: Longest Collatz Sequence", longestCollatzSequence(), 837799},
		{"15: Lattice Paths", latticePaths(), 137846528820},
		{"16: Power Digit Sum", powerDigitSum(), 1366},
		{"17: Number Letter Counts", numberLetterCounts(), 21124},
		{"18: Maximum Path Sum One", maximumPathSumOne(), 1074},
		{"19: Counting Sundays", countingSundays(), 171},
		{"20: Factorial Digit Sum", factorialDigitSum(100), 648},
		{"Bonus: PositionNthPrime(0) test", positionNthPrime(0), 0},
	}
	for _, tC := range testCases {
		if tC.actualResult != tC.expectedResult {
			t.Errorf("Failure. Test case: %v. Actual result: %v. Expected result: %v.", tC.desc, tC.actualResult, tC.expectedResult)
		}
	}
}

func BenchmarkMultiplesOf(b *testing.B) {
	for b.Loop() {
		multiplesOf(3, 5, 1000)
	}
}

func BenchmarkMultiplesOf2(b *testing.B) {
	for b.Loop() {
		multiplesOf2(3, 5, 1000)
	}
}

func BenchmarkEvenFibonacciNumbers(b *testing.B) {
	for b.Loop() {
		evenFibonacciNumbers(1000000)
	}
}

func BenchmarkEvenFibonacciNumbers2(b *testing.B) {
	for b.Loop() {
		evenFibonacciNumbers2(1000000)
	}
}
