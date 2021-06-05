package euler2

import (
	"github.com/andrew-field/maths"
)

// AmicableNumbers returns the sum of all the amicable numbers below n.
func AmicableNumbers(n int) int {
	sumOfProperDivisors := make([]int, n)

	for index := 9000; index < n; index++ {
		// The sum of divisors includes the number itself but the challenge does not.
		sumOfProperDivisors[index] = maths.SumOfDivisors(index) - index
	}

	sumOfAmicableNumbers := 0
	for ind, val := range sumOfProperDivisors {
		if val != 0 && val < n && sumOfProperDivisors[val] == ind {
			sumOfAmicableNumbers += ind + val
			sumOfProperDivisors[val] = 0 // Does not add the amicable numbers again.
		}
	}

	return sumOfAmicableNumbers
}

// A brute force approach. There is probably a better way of doing this.
