package main

import "fmt"

func main() {

	total := 0

	for i := 3; i < 1000; i += 3 {
		total += i
	}

	i := 1
	for j := 5; j < 1000; j += 5 {
		if i == 3 {
			i = 1
			continue
		}
		total += j
		i++
	}

	fmt.Println("Total:", total)

	fmt.Println(45 % 5)
}
