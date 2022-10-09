package euler1

import "time"

// CountingSundays returns the number of Sundays that fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000).
func CountingSundays() int {
	count := 0
	for date := time.Date(1901, 01, 01, 1, 1, 1, 000000000, time.UTC); date.Year() != 2001; date = date.AddDate(0, 1, 0) {
		if date.Weekday() == time.Sunday {
			count++
		}
	}

	return count
}

// "Unfortunately", Go makes this easy with its great time package.
