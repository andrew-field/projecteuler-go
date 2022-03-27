package euler4

import (
	"bufio"
	"io"
	"strconv"

	"github.com/andrew-field/maths"
	"github.com/andrew-field/projecteuler-go/filefunctions"
)

// MaximumPathSumTwo returns the maximum total from top to bottom of a pyramid by starting at the top of the triangle
// and moving to adjacent numbers on the row below.
func MaximumPathSumTwo() int {
	f := filefunctions.OpenFile("p067_triangle.txt")
	defer filefunctions.CloseFile(f)

	reader := bufio.NewReader(f) // This reader is used because of its efficiency with many small reads.

	numbers := make([]int, 0)

	// Read 2 bytes/digits every time to get the right number.
	number := make([]byte, 2)
	for {
		_, err := reader.Read(number)
		if err == io.EOF {
			break
		}
		filefunctions.Check(err)
		value, err := strconv.Atoi(string(number))
		filefunctions.Check(err)
		numbers = append(numbers, value)
	}

	pyramidTree := maths.CreatePyramidTree(numbers...)

	return maths.MaxPath(pyramidTree)
}

// The maths library does most of the work.
