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

	positionOfPrime, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Invalid input. Must be representable as an int.")
	}

	// Make prime and done channel.
	primeChannel := make(chan uint, 100)
	doneChannel := make(chan bool)

	go numbertheory.GetPrimeNumbersContinuously(primeChannel, doneChannel, 100)

	for count := 1; count < positionOfPrime; count++ {
		<-primeChannel
	}
	doneChannel <- true

	fmt.Printf("The prime number at position %v is %v\n", positionOfPrime, <-primeChannel)

	<-doneChannel
}
