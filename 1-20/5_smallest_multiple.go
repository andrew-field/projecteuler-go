package euler1

import "github.com/andrew-field/maths"

// smallestMultiple returns the smallest positive number that is evenly divisible by all of the numbers from 1 to 20.
func smallestMultiple() int {
	// Numbers to 20, not 1.
	numbers := make([]int, 19)
	for ind := range numbers {
		numbers[ind] = ind + 2
	}

	return maths.LCM(numbers...)
}
