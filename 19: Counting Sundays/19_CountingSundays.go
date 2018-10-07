package main

import (
	"fmt"
	"time"
)

func main() {

	// Go is great.
	count := 0
	startDate := time.Date(1901, 01, 01, 1, 1, 1, 000000000, time.UTC)

	for startDate.Year() != 2001 {
		startDate = startDate.AddDate(0, 1, 0)
		if startDate.Weekday() == time.Sunday {
			count++
		}
	}

	fmt.Println("Count", count)
}
