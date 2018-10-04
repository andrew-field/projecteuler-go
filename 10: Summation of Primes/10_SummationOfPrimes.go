package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {

	// Challenge number is 2000000.

	if len(os.Args) != 2 {
		panic("There needs to be 1 arguement, the ceiling for the summation of primes. E.g. 12 for the sum of all the prime numbers under 12.")
	}

	ceiling, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		panic("Invalid input. Must be representable as an uint")
	}

	// Make prime and done channel.
	primeChannel := make(chan uint, 100)
	doneChannel := make(chan bool, 1)

	go numbertheory.GetPrimeNumbersContinuously(primeChannel, doneChannel, 1000)

	var sum uint

	for val := range primeChannel {
		if val < uint(ceiling) {
			sum += val
		} else {
			break
		}
	}

	fmt.Printf("The summation of all the prime numbers under %v is %v\n", ceiling, sum)
	doneChannel <- true
}
