package main

import "fmt"
import "math"

const val = 1000

func main() {
	var a float64 = 1
	var b float64 = 1
	for a < 1000 {
		for b < 1000 {
			if a+b+(math.Hypot(a, b)) == val {
				if math.Floor(math.Hypot(a, b)) == math.Hypot(a, b) {
					c := math.Hypot(a, b)
					fmt.Println("a, b and c:", a, b, c)
					fmt.Println("Product:", int64(a*b*c))
					break
				}
			} else {
				b++
			}
		}
		a++
		b = a
	}
}
