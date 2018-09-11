package main

import "fmt"
import "math"

const val = 1000

func main() {

	var a float64 = 1
	var b float64 = 1

	// Cycle through possible a's and b's.
	for a < 1000 {
		for b < 1000 {
			// This must be true with the two simultaneous equations.
			if c := math.Hypot(a, b); a+b+c == val {
				// As there is only one triplet for which this works, and a and b are 'fine', I don't need to make further checks on c.
				//I.e. it must be the integer in quesion.
				fmt.Println("Product:", int64(a*b*c))
				a = 999
				break
			} else {
				b++
			}
		}
		a++
		b = a
	}
}
