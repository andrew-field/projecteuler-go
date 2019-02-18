package projecteuler1

import (
	"github.com/andrew-field/testing_go/numbertheory"
)

// MaximumPathSumOne returns the maximum total from top to bottom of the following pyramid by starting at the top of the triangle
// and moving to adjacent numbers on the row below.
func MaximumPathSumOne() int {
	// Make the grid for the numbers.
	pyramid := make([][]float64, 15)

	pyramid[0] = []float64{75}
	pyramid[1] = []float64{95, 64}
	pyramid[2] = []float64{17, 47, 82}
	pyramid[3] = []float64{18, 35, 87, 10}
	pyramid[4] = []float64{20, 04, 82, 47, 65}
	pyramid[5] = []float64{19, 01, 23, 75, 03, 34}
	pyramid[6] = []float64{88, 02, 77, 73, 07, 63, 67}
	pyramid[7] = []float64{99, 65, 04, 28, 06, 16, 70, 92}
	pyramid[8] = []float64{41, 41, 26, 56, 83, 40, 80, 70, 33}
	pyramid[9] = []float64{41, 48, 72, 33, 47, 32, 37, 16, 94, 29}
	pyramid[10] = []float64{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14}
	pyramid[11] = []float64{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57}
	pyramid[12] = []float64{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48}
	pyramid[13] = []float64{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31}
	pyramid[14] = []float64{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23}

	return int(numbertheory.GetMaximumPathSumOfPyramidUsingMaximumSlots(pyramid))
}
