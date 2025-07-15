/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

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
