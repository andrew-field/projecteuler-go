package main

import (
	"fmt"
)

func main() {
	max := 0
	answer := 13

	for start := 13; start < 1000000; start++ {
		num := 1
		test := start
		for test != 1 {
			if test%2 == 0 {
				test /= 2
			} else {
				test = 3*test + 1
				num++
				test /= 2
			}
			num++
		}
		if num > max {
			max = num
			answer = start
		}
	}

	fmt.Println("Answer:", answer)
}
