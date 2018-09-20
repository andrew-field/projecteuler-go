package main

import (
	"fmt"
)

func main() {

	// Maximum.
	max := 0
	answer := 0

	for start := 13; start < 1000000; start++ {
		// Number of terms including the starter.
		numberOfTerms := 1

		// Starting number to generate the pattern.
		seq := start
		for seq != 1 {
			if seq%2 == 0 {
				seq /= 2
			} else {
				seq = 3*seq + 1
				numberOfTerms++ // The next term must be even so save some time next.
				seq /= 2
			}
			numberOfTerms++
		}

		// Compare maximum.
		if numberOfTerms > max {
			max = numberOfTerms
			answer = start
		}
	}

	fmt.Println("Answer:", answer)
}
