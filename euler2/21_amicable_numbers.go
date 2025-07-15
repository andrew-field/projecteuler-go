package euler2

import "github.com/andrew-field/maths/v2"

// amicableNumbers returns the sum of all the amicable numbers below |n|.
func amicableNumbers(n int) int {
	n, err := maths.Abs(n)
	if err != nil {
		panic(err)
	}
	sumOfProperDivisors := make(map[int]int)

	for i := range n {
		sum, err := maths.SumOfDivisorsBruteForce(i)
		if err != nil {
			panic(err)
		}
		// The sum of divisors includes the number itself but the challenge does not.
		sum -= i
		if i != sum { // Setup to avoid perfect numbers, which satisfy the condition of d(a) = b and d(b) = a except that a = b.
			sumOfProperDivisors[i] = sum
		}
	}

	sumOfAmicableNumbers := 0
	for key, val := range sumOfProperDivisors {
		if sumOfProperDivisors[val] == key {
			sumOfAmicableNumbers += key + val
			delete(sumOfProperDivisors, val) // Do not add the amicable numbers again.
		}
	}

	return sumOfAmicableNumbers
}

// A brute force approach. There is probably a better way of doing this.
