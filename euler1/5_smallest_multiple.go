/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package euler1

import "github.com/andrew-field/maths/v2"

// smallestMultiple returns the smallest positive number that is evenly divisible by all of the numbers from 1 to 20.
func smallestMultiple() int {
	// Numbers to 20, not 1.
	numbers := make([]int, 19)
	for ind := range numbers {
		numbers[ind] = ind + 2
	}

	lcm, err := maths.LCM(numbers...)
	if err != nil {
		panic(err)
	}

	return lcm
}
