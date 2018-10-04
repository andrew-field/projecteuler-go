package main

import "fmt"

func main() {

	var a float64 = 1

	// Cycle through possible a's and b's.
	for ; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			// This must be true with the two simultaneous equations.
			if 500000 == 1000*(b+a)-b*a {
				// As there is only one triplet for which this works I don't need to make further checks on c.
				fmt.Println("Product:", int(a*b*(1000-b-a)))
				return
			}
		}
	}
}
