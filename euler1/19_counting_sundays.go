/*
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

package euler1

import "time"

// countingSundays returns the number of Sundays that fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000).
func countingSundays() int {
	count := 0
	for date := time.Date(1901, 01, 01, 1, 1, 1, 000000000, time.UTC); date.Year() != 2001; date = date.AddDate(0, 1, 0) {
		if date.Weekday() == time.Sunday {
			count++
		}
	}

	return count
}

// "Unfortunately", Go makes this easy with its great time package.
