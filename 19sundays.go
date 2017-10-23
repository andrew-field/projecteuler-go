package main

import (
	"fmt"
	"time"
)

func main() {

	// Go is great.
	count := 0
	date := time.Date(1901, 01, 06, 1, 1, 1, 000000000, time.UTC)
	for date.Year() != 2001 {
		date = date.AddDate(0, 0, 7)
		if date.Day() == 1 {
			count++
		}
	}
	fmt.Println("Count", count)
}
