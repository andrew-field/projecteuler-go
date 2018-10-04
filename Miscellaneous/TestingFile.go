package main

import (
	"fmt"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {
	for i := 2; i < 10; i++ {
		primeFactorChannel := make(chan uint, 100)
		go numbertheory.GetPrimeFactorisation(primeFactorChannel, uint(i))
		fmt.Println("i:", i)
		for val := range primeFactorChannel {
			fmt.Println(val)
		}
	}
}
