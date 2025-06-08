package euler1

import (
	"testing"

	"github.com/andrew-field/projecteuler-go/test"
)

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []test.TestCase{
		{Desc: "1a: Multiples of 3 and 5", ActualResult: multiplesOf(3, 5, 1000), ExpectedResult: 233168},
		{Desc: "1b: Multiples of 3 and 5", ActualResult: multiplesOf2(3, 5, 1000), ExpectedResult: 233168},
		{Desc: "2a: Even Fibonacci Numbers", ActualResult: evenFibonacciNumbers(4000000), ExpectedResult: 4613732},
		{Desc: "2b: Even Fibonacci Numbers", ActualResult: evenFibonacciNumbers2(4000000), ExpectedResult: 4613732},
		{Desc: "3: Largest Prime Factor", ActualResult: largestPrimeFactor(600851475143), ExpectedResult: 6857},
		{Desc: "4: Largest Palindrome Product", ActualResult: largestPalindromeProduct(), ExpectedResult: 906609},
		{Desc: "5: Smallest Multiple", ActualResult: smallestMultiple(), ExpectedResult: 232792560},
		{Desc: "6: Sum Square Difference", ActualResult: sumSquareDifference(1, 100), ExpectedResult: 25164150},
		{Desc: "7: 10001st Prime Number", ActualResult: positionNthPrime(10001), ExpectedResult: 104743},
		{Desc: "8: Largest Product in a Series", ActualResult: largestProductInASeries(), ExpectedResult: 23514624000},
		{Desc: "9: Special Pythagorean triplet", ActualResult: specialPythagoreanTriplet(), ExpectedResult: 31875000},
		{Desc: "10: Summation of Primes", ActualResult: summationOfPrimes(2000000), ExpectedResult: 142913828922},
		{Desc: "11: Largest Product in a Grid", ActualResult: largestProductInAGrid(), ExpectedResult: 70600674},
		{Desc: "12: Highly Divisible Triangular Number", ActualResult: highlyDivisibleTriangularNumber(500), ExpectedResult: 76576500},
		{Desc: "13: Large Sum", ActualResult: largeSum(), ExpectedResult: 5537376230},
		{Desc: "14: Longest Collatz Sequence", ActualResult: longestCollatzSequence(), ExpectedResult: 837799},
		{Desc: "15: Lattice Paths", ActualResult: latticePaths(), ExpectedResult: 137846528820},
		{Desc: "16: Power Digit Sum", ActualResult: powerDigitSum(), ExpectedResult: 1366},
		{Desc: "17: Number Letter Counts", ActualResult: numberLetterCounts(), ExpectedResult: 21124},
		{Desc: "18: Maximum Path Sum One", ActualResult: maximumPathSumOne(), ExpectedResult: 1074},
		{Desc: "19: Counting Sundays", ActualResult: countingSundays(), ExpectedResult: 171},
		{Desc: "20: Factorial Digit Sum", ActualResult: factorialDigitSum(100), ExpectedResult: 648},
	}

	test.CheckResults(testCases, t)
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
