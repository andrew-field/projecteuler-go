package euler2

import "github.com/andrew-field/maths/v2"

// amicableNumbers returns the sum of all the amicable numbers below n.
func amicableNumbers(n int) int {
	sumOfProperDivisors := make([]int, n)

	for index := 2; index < n; index++ {
		sum, err := maths.SumOfDivisorsBruteForce(index)
		if err != nil {
			panic(err)
		}
		// The sum of divisors includes the number itself but the challenge does not.
		sumOfProperDivisors[index] = sum - index
		if index == sumOfProperDivisors[index] { // Setup to avoid perfect numbers, which satisfy the condition of d(a) = b and d(b) = a except that a = b.
			sumOfProperDivisors[index] = 1
		}
	}
	sumOfProperDivisors[0], sumOfProperDivisors[1] = 1, 1

	sumOfAmicableNumbers := 0
	for ind, val := range sumOfProperDivisors {
		if val != 1 && val < n && sumOfProperDivisors[val] == ind {
			sumOfAmicableNumbers += ind + val
			sumOfProperDivisors[val] = 1 // Do not add the amicable numbers again.
		}
	}

	return sumOfAmicableNumbers
}

// A brute force approach. There is probably a better way of doing this.
