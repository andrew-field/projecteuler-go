/*
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

3
7 4
2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in triangle.txt (right click and 'Save Link/Target As...'), a 15K text file containing a triangle with one-hundred rows.

NOTE: This is a much more difficult version of Problem 18.
*/

package euler4

import (
	"strconv"

	_ "embed"

	"github.com/andrew-field/maths/v2"
)

//go:embed p067_triangle.txt
var s string

// maximumPathSumTwo returns the maximum total from top to bottom of a pyramid by starting at the top of the triangle
// and moving to adjacent numbers on the row below.
func maximumPathSumTwo() int {
	numbers := make([]int, len(s)/2)

	for i := 0; i < len(s); i += 2 {
		value, err := strconv.Atoi(s[i : i+2])
		if err != nil {
			panic(err)
		}

		numbers[i/2] = value
	}

	return maths.MaxPath(maths.CreatePyramidTree(numbers...))
}

// The maths library does most of the work.
