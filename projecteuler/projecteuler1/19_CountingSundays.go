package projecteuler1

import (
	"time"
)

// CountingSundays returns the number of Sundays that fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000).
func CountingSundays() int {
	// Go is great, makes it too easy :).
	count := 0
	for startDate := time.Date(1901, 01, 01, 1, 1, 1, 000000000, time.UTC); startDate.Year() != 2001; startDate = startDate.AddDate(0, 1, 0) {
		if startDate.Weekday() == time.Sunday {
			count++
		}
	}

	return count
}
