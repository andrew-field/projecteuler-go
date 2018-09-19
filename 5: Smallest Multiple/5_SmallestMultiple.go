package main

import (
	"fmt"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {
	// Numbers to 20, not 1.
	numbers := make([]uint, 19)
	for ind := range numbers {
		numbers[ind] = uint(ind) + 2
	}

	fmt.Println("Total:", numbertheory.LowestCommonMultiple(numbers...))
}
