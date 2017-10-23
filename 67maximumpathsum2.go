package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Height of pyramid.
	length := 100

	// Make the grid for the numbers.
	grid := make([][]float64, length)
	for ind := range grid {
		grid[ind] = make([]float64, ind+1)
	}

	// Open file.
	f, err := os.Open("/home/andrew/go/src/github.com/andrew-field/go-testing/p067_triangle.txt")
	check(err)

	// Read 2 digits every time.
	read := make([]byte, 2)
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			_, err = f.Read(read)
			check(err)
			grid[i][j], err = strconv.ParseFloat(string(read), 64)
			check(err)
		}
	}

	// Make the slice to hold the individual maximum numbers at each index of grid.
	max := make([][]float64, length)

	// Maximum.
	var maxi float64

	// Go through each index starting at the top. Populate the max array with the maximum possible sum at each index by adding the
	// grid number to the maximum of the two maximums for the above indexes. Left and right hand indexes are taken account of.
	for ind := range grid {
		max[ind] = make([]float64, ind+1)
		for ind2 := range grid[ind] {
			if ind == 0 {
				max[ind][ind2] = grid[ind][ind2]
			} else if ind2 == 0 {
				max[ind][ind2] = grid[ind][ind2] + max[ind-1][0]
			} else if ind == ind2 {
				max[ind][ind2] = grid[ind][ind2] + max[ind-1][ind2-1]
			} else {
				max[ind][ind2] = grid[ind][ind2] + math.Max(max[ind-1][ind2-1], max[ind-1][ind2])
			}
			// If on the last row, find the maximum.
			if ind == length-1 {
				if max[length-1][ind2] > maxi {
					maxi = max[length-1][ind2]
				}
			}
		}
	}
	fmt.Println("Answer", maxi)
}
