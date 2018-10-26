package numbertheory

import (
	"math"
)

// GetMaximumPathSumOfPyramidUsingMaximumSlots returns the maximum sum of every path from top to bottom of a pyramid while
// moving only to adjacent numbers on the row below.
func GetMaximumPathSumOfPyramidUsingMaximumSlots(pyramid [][]float64) float64 {
	// Height of the pyramid/Length of longest row.
	length := len(pyramid)

	// Make the grid for the numbers.
	pyramidWithMaximumSlots := make([][][]float64, length)

	// Each index of the pyramid has a corresponding "max slot" (The zeroes).
	for ind := range pyramidWithMaximumSlots {
		pyramidWithMaximumSlots[ind] = make([][]float64, 0)
		for j := 0; j <= ind; j++ {
			pyramidWithMaximumSlots[ind] = append(pyramidWithMaximumSlots[ind], []float64{pyramid[ind][j], 0})
		}
	}

	// Go through each index starting at the top.
	// Populate each max slot of each index with the maximum possible sum with which it is possible
	// to reach that index. Do this by adding the grid number at the index to the maximum of the
	// two max slots for the above indexes (Directly above or left).
	for ind := range pyramidWithMaximumSlots {
		// The top.
		if ind == 0 {
			pyramidWithMaximumSlots[ind][0][1] = pyramidWithMaximumSlots[ind][0][0]
			continue
		}
		for ind2 := range pyramidWithMaximumSlots[ind] {
			if ind2 == 0 { // The leftmost indexes.
				pyramidWithMaximumSlots[ind][ind2][1] = pyramidWithMaximumSlots[ind][ind2][0] + pyramidWithMaximumSlots[ind-1][ind2][1]
			} else if ind == ind2 { // The rightmost indexes.
				pyramidWithMaximumSlots[ind][ind2][1] = pyramidWithMaximumSlots[ind][ind2][0] + pyramidWithMaximumSlots[ind-1][ind2-1][1]
			} else { // The rest.
				pyramidWithMaximumSlots[ind][ind2][1] = pyramidWithMaximumSlots[ind][ind2][0] + math.Max(pyramidWithMaximumSlots[ind-1][ind2-1][1], pyramidWithMaximumSlots[ind-1][ind2][1])
			}
		}
	}

	// Maximum.
	var max float64

	// Find the maximum of the max slots in the final row.
	for _, val := range pyramidWithMaximumSlots[length-1] {
		max = math.Max(max, val[1])
	}

	return max
}

var pyramidUsingRecursiveFunction [][]float64

// GetMaximumPathSumOfPyramidUsingRecursiveFunction returns the maximum sum of every path from top to bottom of a pyramid while
// moving only to adjacent numbers on the row below. This method uses a recursive function and is relatively slow.
func GetMaximumPathSumOfPyramidUsingRecursiveFunction(pyramid [][]float64) float64 {
	// Maximum.
	var max float64
	length := len(pyramid)
	pyramidUsingRecursiveFunction = pyramid

	// Find the maximum of the maximums in the final row.
	for i := 0; i < length; i++ {
		max = math.Max(max, getMax(length-1, i))
	}

	return max
}

func getMax(x, y int) float64 {
	// The top.
	if x == 0 {
		return pyramidUsingRecursiveFunction[x][y]
	}
	// The leftmost indexes.
	if y == 0 {
		return pyramidUsingRecursiveFunction[x][y] + getMax(x-1, y)
	}
	// The rightmost indexes.
	if y == x {
		return pyramidUsingRecursiveFunction[x][y] + getMax(x-1, y-1)
	}
	// The rest.
	return pyramidUsingRecursiveFunction[x][y] + math.Max(getMax(x-1, y-1), getMax(x-1, y))
}
