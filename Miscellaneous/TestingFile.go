package main

import (
	"fmt"

	"github.com/andrew-field/testing_go/numbertheory"
)

func main() {
	sdf := numbertheory.GetDivisorsOfANumberInASlice(284)

	fmt.Println(sdf)
}
