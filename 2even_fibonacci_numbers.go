package main

import "fmt"

func main() {

	num1 := 1
	num2 := 1
	total := 0

	for num1 < 4000000 {
		num2 += num1
		if num2%2 == 0 {
			total += num2
		}
		num1, num2 = num2, num1
	}

	fmt.Println(total)
}
