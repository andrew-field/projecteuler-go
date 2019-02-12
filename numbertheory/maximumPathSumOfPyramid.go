package numbertheory

import (
	"math"
)

// GetMaximumPathSumOfPyramidUsingMaximumSlots returns the maximum of every path sum from top to bottom of a pyramid while
// moving only to adjacent numbers on the row below.
func GetMaximumPathSumOfPyramidUsingMaximumSlots(pyramid [][]float64) float64 {
	if pyramid == nil {
		panic("The pyramid can not be nil.")
	}

	// Height of the pyramid/Length of longest row.
	length := len(pyramid)

	// Safety checking of correctly constructed pyramid as a 2D slice.
	if length == 0 {
		panic("The pyramid can not have zero length.")
	}
	for ind := range pyramid {
		if len(pyramid[ind]) != ind+1 {
			panic("The pyramid is not properly constructed.")
		}
	}

	// The Coordinate can correspond to each index of the 2D slice.
	type Coordinate struct {
		outerIndex, innerIndex int
	}

	// Make the map to store for each index of the pyramid the maximum possible sum with which it is possible
	// to reach that index.
	maximumMap := make(map[Coordinate]float64)

	// The top is obvious.
	maximumMap[Coordinate{0, 0}] = pyramid[0][0]

	// Go through each index excluding the top.
	// Populate th maximumMap including each index of the pyramid. Do this by adding the value at each index of the pyramid to the maximum of the
	// two values stored in the maximumMap for the indexes above (directly above and, above and left).
	for outerIndex := 1; outerIndex < length; outerIndex++ {
		for innerIndex := range pyramid[outerIndex] {
			if innerIndex == 0 { // The leftmost indexes.
				maximumMap[Coordinate{outerIndex, innerIndex}] = pyramid[outerIndex][innerIndex] + maximumMap[Coordinate{outerIndex - 1, innerIndex}]
			} else if outerIndex == innerIndex { // The rightmost indexes.
				maximumMap[Coordinate{outerIndex, innerIndex}] = pyramid[outerIndex][innerIndex] + maximumMap[Coordinate{outerIndex - 1, innerIndex - 1}]
			} else { // The rest.
				maximumMap[Coordinate{outerIndex, innerIndex}] = pyramid[outerIndex][innerIndex] + math.Max(maximumMap[Coordinate{outerIndex - 1, innerIndex - 1}], maximumMap[Coordinate{outerIndex - 1, innerIndex}])
			}
		}
	}

	// Maximum.
	var max float64

	// Find the maximum in the maximumMap at the indexes of the bottom and final row.
	for index := 0; index < length; index++ {
		max = math.Max(max, maximumMap[Coordinate{length - 1, index}])
	}

	return max
}

// GetMaximumPathSumOfPyramidUsingRecursiveFunction returns the maximum of every path sum from top to bottom of a pyramid while
// moving only to adjacent numbers on the row below.
func GetMaximumPathSumOfPyramidUsingRecursiveFunction(pyramid [][]float64) float64 {
	if pyramid == nil {
		panic("The pyramid can not be nil.")
	}

	// Height of the pyramid/Length of longest row cannot be 0.
	if len(pyramid) == 0 {
		panic("The pyramid can not have zero length.")
	}
	// Safety checking of correctly constructed pyramid as a 2D slice.
	for ind := range pyramid {
		if len(pyramid[ind]) != ind+1 {
			panic("The pyramid is not properly constructed.")
		}
	}

	return getMax(pyramid)
}

func getMax(pyramid [][]float64) float64 {
	// A single number, the bottom of the pyramid.
	if len(pyramid) == 1 {
		return pyramid[0][0]
	}

	// Because of how go uses memory for slices, a copy is needed to have to distinct slices not altering each other.
	// See https://blog.golang.org/go-slices-usage-and-internals for more.
	copyOfPyramid := make([][]float64, len(pyramid))
	copy(copyOfPyramid, pyramid)

	return pyramid[0][0] + math.Max(getMax(getLowerPyramid(pyramid, "Left")), getMax(getLowerPyramid(copyOfPyramid, "Right")))
}

func getLowerPyramid(pyramid [][]float64, position string) [][]float64 {
	pyramid = pyramid[1:]

	for ind := range pyramid {
		if position == "Left" {
			pyramid[ind] = pyramid[ind][:len(pyramid[ind])-1]
		} else if position == "Right" {
			pyramid[ind] = pyramid[ind][1:]
		} else {
			panic("Encountered unexpected position.")
		}
	}

	return pyramid
}
