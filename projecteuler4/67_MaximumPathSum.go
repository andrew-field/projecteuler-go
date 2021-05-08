package projecteuler4

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"

	"github.com/andrew-field/testing_go/numbertheory"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// MaximumPathSumTwo returns the maximum total from top to bottom of the following pyramid by starting at the top of the triangle
// and moving to adjacent numbers on the row below.
func MaximumPathSumTwo() int {
	length := 100

	// Make the grid for the numbers.
	pyramid := make([][]float64, length)

	// Open file.
	absPath, err := filepath.Abs("p067_triangle.txt")
	check(err)
	f, err := os.Open(absPath)
	check(err)
	defer func() {
		f.Close()
	}()
	reader := bufio.NewReader(f)

	// Read 2 digits every time and populate the pyramid.
	number := make([]byte, 2)
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			_, err = reader.Read(number)
			check(err)
			value, err := strconv.ParseFloat(string(number), 64)
			check(err)
			pyramid[i] = append(pyramid[i], value)
		}
	}

	return int(numbertheory.GetMaximumPathSumOfPyramidUsingMaximumSlots(pyramid))
}
