package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {

	// Challenge position is 10001st.

	if len(os.Args) != 2 {
		panic("There needs to be 1 arguement, the position of the prime number you would like to know. E.g. \"12\" for 12th.")
	}

	positionOfPrime, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		panic("Invalid input. Must be representable as an uint.")
	}

	// Make prime and done channel.
	primeChannel, doneChannel := numbertheory.GetPrimeNumbersContinuously(100)

	var count uint64 = 1
	var prime uint
	for ; count <= positionOfPrime; count++ {
		prime = <-primeChannel
	}
	doneChannel <- true

	fmt.Printf("The prime number at position %v is %v\n", positionOfPrime, prime)

	<-doneChannel
}
