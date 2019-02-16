package numbertheory

import "sort"

func channelIntoSlice(function func(uint) chan uint, number uint) []uint {
	channel := function(number)
	allValues := make([]uint, 0)
	for val := range channel {
		allValues = append(allValues, val)
	}
	sort.Slice(allValues, func(i, j int) bool { return allValues[i] < allValues[j] })
	return allValues
}
