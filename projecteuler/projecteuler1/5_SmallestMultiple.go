package projecteuler1

import (
	"github.com/andrew-field/testing_go/numbertheory"
)

// SmallestMultiple returns the smallest positive number that is evenly divisible by all of the numbers from 1 to 20.
func SmallestMultiple() int {
	// Numbers to 20, not 1.
	numbers := make([]uint, 19)
	for ind := range numbers {
		numbers[ind] = uint(ind) + 2
	}

	return int(numbertheory.LowestCommonMultiple(numbers...))
}
