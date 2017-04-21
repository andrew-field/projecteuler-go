package main

import "fmt"

func main() {

	num1 := 1
	num2 := 1
	buffer := 0
	total := 0

	for num2 < 4000000 {
		if num2%2 == 0 {
			total += num2
		}
		buffer = num2
		num2 += num1
		num1 = buffer
	}

	fmt.Println(total)
}
