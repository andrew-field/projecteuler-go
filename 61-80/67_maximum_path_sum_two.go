package euler4

import (
	"bufio"
	"io"
	"os"
	"strconv"

	"github.com/andrew-field/maths"
)

// maximumPathSumTwo returns the maximum total from top to bottom of a pyramid by starting at the top of the triangle
// and moving to adjacent numbers on the row below.
func maximumPathSumTwo() int {
	f, err := os.Open("p067_triangle.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f) // This reader is used because of its efficiency with many small reads.

	numbers := make([]int, 0)

	// Read 2 bytes/digits every time to get the right number.
	number := make([]byte, 2)
	for {
		_, err := reader.Read(number)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		value, err := strconv.Atoi(string(number))
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, value)
	}

	pyramidTree := maths.CreatePyramidTree(numbers...)

	return maths.MaxPath(pyramidTree)
}

// The maths library does most of the work.
