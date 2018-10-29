package main

import (
	"fmt"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {
	sdf := numbertheory.GetDivisorsOfANumberBruteForce(284)

	for val := range sdf {
		fmt.Println(val)
	}
}
