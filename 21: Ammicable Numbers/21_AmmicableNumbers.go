package main

import (
	"fmt"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {
	sumOfProperDivisors := make([]uint, 10000)

	var index uint = 2
	for ; index < 10000; index++ {
		divisorChannel := numbertheory.GetDivisorsOfANumber(index)

		// The divisors includes the number itself but the challenge does not.
		sum := -index
		for val := range divisorChannel {
			sum += val
		}

		sumOfProperDivisors[index] = sum
	}

	var sumOfAmicableNumbers uint
	for ind, val := range sumOfProperDivisors {
		if val != 1 && val != 0 && uint(ind) != val && val < 10000 && sumOfProperDivisors[val] == uint(ind) {
			sumOfAmicableNumbers += uint(ind) + val
			sumOfProperDivisors[val] = 1
			sumOfProperDivisors[ind] = 1
		}
	}

	fmt.Println("Answer:", sumOfAmicableNumbers)
}
