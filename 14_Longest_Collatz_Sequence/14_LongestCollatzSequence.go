package main

import (
	"fmt"
)

func main() {

	// Maximum.
	maxNumberOfTerms := 0
	// Answer (Starting number).
	answer := 0

	// Grid of numbers to reduce the number of sequences generated.
	gridOfNumbers := make([]int, 1000001)
	for ind := range gridOfNumbers {
		gridOfNumbers[ind] = ind
	}

	for index := 1000000; index > 0; index-- {
		// Starter numbers which have been part of previous sequences will form sequences
		// which must be shorter than the previous maximum.
		if gridOfNumbers[index] == 1 {
			continue
		}

		// Number of terms including the starter.
		numberOfTerms := 1

		// Generate the sequence. The grid reference for numbers along the way will be set to 1
		// to indicate sequences which should be skipped later.
		for gridOfNumbers[index] != 1 {
			if gridOfNumbers[index]%2 == 0 {
				gridOfNumbers[index] /= 2
			} else {
				gridOfNumbers[index] = 3*gridOfNumbers[index] + 1 // The next term must be even so save some time next.
				if gridOfNumbers[index] <= 1000000 && gridOfNumbers[gridOfNumbers[index]] != 1 {
					gridOfNumbers[gridOfNumbers[index]] = 1
				}
				numberOfTerms++
				gridOfNumbers[index] /= 2
			}
			if gridOfNumbers[index] <= 1000000 && gridOfNumbers[gridOfNumbers[index]] != 1 {
				gridOfNumbers[gridOfNumbers[index]] = 1
			}
			numberOfTerms++
		}

		// Compare maximum.
		if numberOfTerms > maxNumberOfTerms {
			maxNumberOfTerms = numberOfTerms
			answer = index
		}
	}

	fmt.Println("Answer:", answer)
}
