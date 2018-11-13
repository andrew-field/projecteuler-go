package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {

	// Challenge number is 600851475143.

	if len(os.Args) != 2 {
		panic("There needs to be 1 arguement, the number for which you would like to know the largest prime factor. E.g. 6542.")
	}

	numberToFactorise, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		panic("Invalid input. Must be representable as an uint")
	}

	// Make factor channel.
	primeFactorChannel := numbertheory.GetPrimeFactorisation(uint(numberToFactorise))

	var largestPrimeFactor uint

	for val := range primeFactorChannel {
		largestPrimeFactor = val
	}

	fmt.Println("Largest prime factor:", largestPrimeFactor)
}
