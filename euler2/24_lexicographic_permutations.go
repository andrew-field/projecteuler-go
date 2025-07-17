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

	fact, err := maths.Factorial(10)
	if err != nil {
		panic(err)
	}
	for i := 9; n != 0; i-- {
		fact /= i + 1
		position := n / fact
		if n%fact == 0 {
			position--
			n -= fact
		}
		permutation[9-i] = numbers[position]
		numbers = slices.Delete(numbers, position, position+1)
		n -= position * fact
	}

	for i, v := range numbers {
		permutation[9-i] = v
	}

	return permutation[0]*1000000000 +
		permutation[1]*100000000 +
		permutation[2]*10000000 +
		permutation[3]*1000000 +
		permutation[4]*100000 +
		permutation[5]*10000 +
		permutation[6]*1000 +
		permutation[7]*100 +
		permutation[8]*10 +
		permutation[9]
}
