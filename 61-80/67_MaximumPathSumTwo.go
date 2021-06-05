package euler4

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/andrew-field/maths"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// MaximumPathSumTwo returns the maximum total from top to bottom of a pyramid by starting at the top of the triangle
// and moving to adjacent numbers on the row below.
func MaximumPathSumTwo() int {
	// Open file.
	absPath, err := filepath.Abs("p067_triangle.txt")
	check(err)
	f, err := os.Open(absPath)
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f) // This reader is used because of its efficiency with many small reads.

	numbers := make([]int, 0)

	// Read 2 bytes/digits every time to get the right number.
	number := make([]byte, 2)
	for {
		_, err = reader.Read(number)
		if err == io.EOF {
			break
		}
		check(err)
		value, err := strconv.Atoi(string(number))
		check(err)
		numbers = append(numbers, value)
	}

	pyramidTree := maths.CreatePyramidTree(numbers...)

	return maths.MaxPath(pyramidTree)
}

// The maths library does most of the work.
