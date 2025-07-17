/*
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/

package euler2

import (
	"github.com/andrew-field/maths/v2"
)

func nonAbundantSums() int {
	// We know 12 is the first abundant number.
	abundantNumbers := []int{12}

	// 28122 is the largest number we need to check to be the sum of two abundant numbers.
	// As 12 is the smallest abundant number, the largest number we need to check for being abundant is 28110.
	for i := 13; i < 28111; i++ {
		sumOfDivisors, err := maths.SumOfDivisors(i)
		if err != nil {
			panic(err)
		}
		if sumOfDivisors-i > i {
			abundantNumbers = append(abundantNumbers, i)
		}
	}

	// We will calculate the final sum by subtracting the sum of all positive integers which can be written as the sum of two abundant numbers up to 28123 from the sum of all positive integers up to 28123.
	// Sum of all integers from 1 to 28122, using the formula n(n+1)/2.
	totalSum := 28122 * 28123 / 2

	abundantSums := map[int]struct{}{}

	// Calculate all the possible sums of two abundant numbers, less than 28123 and subtract them from the total sum.
	for i := 0; i < len(abundantNumbers); i++ {
		for j := i; j < len(abundantNumbers); j++ {
			if abundantNumbers[i]+abundantNumbers[j] < 28123 {
				sum := abundantNumbers[i] + abundantNumbers[j]
				if _, ok := abundantSums[sum]; !ok {
					abundantSums[sum] = struct{}{}
					totalSum -= sum
				}
			}
		}
	}

	return totalSum
}
