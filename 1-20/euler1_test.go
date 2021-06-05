package euler1

import "testing"

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []struct {
		desc           string
		actualResult   int
		expectedResult int
	}{
		{"1: Multiples of 3 and 5", MultiplesOf(3, 5, 1000), 233168},
		{"2: Even Fibonacci Numbers", EvenFibonacciNumbers(4000000), 4613732},
		{"3: Largest Prime Factor", LargestPrimeFactor(600851475143), 6857},
		{"4: Largest Palindrome Product", LargestPalindromeProduct(), 906609},
		{"5: Smallest Multiple", SmallestMultiple(), 232792560},
		{"6: Sum Square Difference", SumSquareDifference(1, 100), 25164150},
		{"7: 10001st Prime Number", PositionNthPrime(10001), 104743},
		{"8: Largest Product in a Series", LargestProductInASeries(), 23514624000},
		{"9: Special Pythagorean triplet", SpecialPythagoreanTriplet(), 31875000},
		{"10: Summation of Primes", SummationOfPrimes(2000000), 142913828922},
		{"11: Largest Product in a Grid", LargestProductInAGrid(), 70600674},
		//{"12: Highly Divisible Triangular Number", HighlyDivisibleTriangularNumber(500), 76576500},
		{"13: Large Sum", LargeSum(), 5537376230},
		{"14: Longest Collatz Sequence", LongestCollatzSequence(), 837799},
		{"15: Lattice Paths", LatticePaths(), 137846528820},
		{"16: Power Digit Sum", PowerDigitSum(), 1366},
		{"17: Number Letter Counts", NumberLetterCounts(), 21124},
		{"18: Maximum Path Sum One", MaximumPathSumOne(), 1074},
		{"19: Counting Sundays", CountingSundays(), 171},
		{"20: Factorial Digit Sum", FactorialDigitSum(100), 648},
	}
	for _, tC := range testCases {
		if tC.actualResult != tC.expectedResult {
			t.Errorf("Failure. Test case: %v. Actual result: %v. Expected result: %v.", tC.desc, tC.actualResult, tC.expectedResult)
		}
	}
}
