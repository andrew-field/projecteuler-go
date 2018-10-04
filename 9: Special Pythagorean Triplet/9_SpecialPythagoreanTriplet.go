package main

import "fmt"
import "math"

func main() {

	var a float64 = 1

	// Cycle through possible a's and b's.
	for ; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			// This must be true with the two simultaneous equations.
			if c := math.Hypot(a, b); a+b+c == 1000 {
				// As there is only one triplet for which this works, and a and b are 'fine', I don't need to make further checks on c.
				// It must be the integer in quesion.
				fmt.Println(c)
				fmt.Println("Product:", int(a*b*c))
				return
			}
		}
	}
}
