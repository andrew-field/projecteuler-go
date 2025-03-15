package euler1

import (
	"github.com/andrew-field/maths/v2"
)

// maximumPathSumOne returns the maximum total from top to bottom of a pyramid by starting at the top of the triangle
// and moving to adjacent numbers on the row below.
func maximumPathSumOne() int {

	// Written on multiple lines to make the pyramid result more clear.
	pyramid := []int{75}
	pyramid = append(pyramid, 95, 64)
	pyramid = append(pyramid, 17, 47, 82)
	pyramid = append(pyramid, 18, 35, 87, 10)
	pyramid = append(pyramid, 20, 04, 82, 47, 65)
	pyramid = append(pyramid, 19, 01, 23, 75, 03, 34)
	pyramid = append(pyramid, 88, 02, 77, 73, 07, 63, 67)
	pyramid = append(pyramid, 99, 65, 04, 28, 06, 16, 70, 92)
	pyramid = append(pyramid, 41, 41, 26, 56, 83, 40, 80, 70, 33)
	pyramid = append(pyramid, 41, 48, 72, 33, 47, 32, 37, 16, 94, 29)
	pyramid = append(pyramid, 53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14)
	pyramid = append(pyramid, 70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57)
	pyramid = append(pyramid, 91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48)
	pyramid = append(pyramid, 63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31)
	pyramid = append(pyramid, 04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23)

	pyramidTree := maths.CreatePyramidTree(pyramid...)

	return maths.MaxPath(pyramidTree)
}

// The maths library does most of the work.
