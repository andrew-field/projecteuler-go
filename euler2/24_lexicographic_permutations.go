/*
A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

package euler2

import (
	"slices"

	"github.com/andrew-field/maths/v2"
)

func lexicographicPermutations(n int) int {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	permutation := make([]int, 10)

	fact, err := maths.Factorial(9)
	if err != nil {
		panic(err)
	}
	for i := 9; ; i-- {
		// When inspecting the position of the current first digit in numbers, there are x number of permutations of the remaining digits, where x is the factorial of the number of remaining digits.
		// We can determine how many complete sets of x permutations fit into the remaining n permutations.
		// For each complete set with at least one remaining permutation, we must increment the index value of the next digit in the final permutation. The numbers are ordered, so the digit in the following position is also the next largest.
		// As n and fact are ints, position equals floor(n / fact) i.e. the number of complete sets with at least one remaining permutation (if n is not divisible by fact).
		// If position equals zero, the desired permutation is found within the possible permutations of the remaining digits and the next digit in the final permutation must be the current
		// smallest digit remaining numbers, which is at index 0 as they are ordered.
		position := n / fact

		// If y = n / fact and fact | n, then there are a y-1 number of complete sets of permutations in n with at least one permutation remaining.
		// Furthermore in this case, we know the remaining part of the final permutation must be the digit in y-1 position followed by the last permutation of the remaining digits in lexicographical order i.e. the remaining numbers reversed.
		// This will always be true eventually.
		if n%fact == 0 {
			position--
			permutation[9-i] = numbers[position]
			numbers = slices.Delete(numbers, position, position+1)

			// The remaining digits (perhaps only digit) left in the permutation are the remaining numbers in reverse order.
			for i, v := range numbers {
				permutation[9-i] = v
			}

			result, err := maths.DigitsToInt(permutation...)
			if err != nil {
				panic(err)
			}

			return result
		}
		permutation[9-i] = numbers[position]
		numbers = slices.Delete(numbers, position, position+1)

		// The next position * fact permutations have been accounted for, so we can reduce n by position * fact.
		n -= position * fact

		// Reduce the factorial for the next digit.
		fact /= i
	}
}
