package main

import (
	"fmt"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {

	// // Old method.
	// // Primes below 20.
	// primes := []int{2, 3, 5, 7, 11, 13, 17, 19}

	// Old and current method.
	// Numbers to 20, not 1.
	numbers := make([]uint, 19)
	for ind := range numbers {
		numbers[ind] = uint(ind) + 2
	}

	fmt.Println("Total:", numbertheory.LowestCommonMultiple(numbers...))

	// // Old Method.
	// var total float64 = 1

	// // lcm by raising each prime number to the maximum number of times
	// // that prime number appears in any of the numbers and then multiplying each of
	// // these results.
	// for _, prime := range primes {
	// 	count := 0
	// 	for _, val := range num {
	// 		temp := 0
	// 		for ; val%prime == 0; temp++ {
	// 			val /= prime
	// 		}
	// 		if temp > count {
	// 			count = temp
	// 		}
	// 	}

	// 	// Multiply to total.
	// 	total *= math.Pow(float64(prime), float64(count))
	// }

	// fmt.Println("Total:", int(total))
}
