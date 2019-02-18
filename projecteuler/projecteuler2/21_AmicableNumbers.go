package projecteuler2

import (
	"github.com/andrew-field/testing_go/numbertheory"
)

// AmicableNumbers returns the sum of all the amicable numbers under 10000.
func AmicableNumbers() int {
	sumOfProperDivisors := make([]uint, 10000)

	var index uint = 2
	for ; index < 10000; index++ {
		divisorChannel := numbertheory.GetDivisorsOfANumber(index)

		// The divisors includes the number itself but the challenge does not.
		sum := -index
		for val := range divisorChannel {
			sum += val
		}

		sumOfProperDivisors[index] = sum
	}

	sumOfAmicableNumbers := 0
	for ind, val := range sumOfProperDivisors {
		if val != 1 && val != 0 && uint(ind) != val && val < 10000 && sumOfProperDivisors[val] == uint(ind) {
			sumOfAmicableNumbers += ind + int(val)
			sumOfProperDivisors[val], sumOfProperDivisors[ind] = 1, 1
		}
	}

	return sumOfAmicableNumbers
}
