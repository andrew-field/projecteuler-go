// Quicker by hand (lcm).

package main

import "fmt"

func main() {

	// Primes below 20.
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}

	// Numbers below 20, not 1.
	num := make([]int, 19)
	for ind := range num {
		num[ind] = ind + 2
	}

	total := 1

	// lcm by raising each prime number to the maximum number of times
	// that prime appears in any of the nums and then multiplying each of
	// these results.
	for _, prime := range primes {
		count := 0
		for _, val := range num {
			temp := 0
			for val%prime == 0 {
				temp++
				val /= prime
			}
			if temp > count {
				count = temp
			}
		}
		// val becomes prime rasied to the power count.
		mult := 1
		for count > 0 {
			mult *= prime
			count--
		}

		// Multiply to total.
		total *= mult
	}

	fmt.Println("Total:", total)

}
