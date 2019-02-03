package main

import (
	"fmt"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {
	for val := range numbertheory.GetAllPrimeNumbersBelowCeiling(100) {
		fmt.Println(val)
	}
}
