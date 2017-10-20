package main

import (
	"bufio"
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
	// Make the grid for the numbers.
	grid := make([][]float64, 15)
	for ind := range grid {
		grid[ind] = make([]float64, 0)
	}

	f, err := os.Open("/home/andrew/go/src/github.com/andrew-field/go-testing/p067_triangle.txt")
	check(err)

	r1 := bufio.NewReader(f)
	b1, err := r1.Peek(2)
	check(err)
	fmt.Println("String", string(b1))
	test, err := strconv.Atoi(string(b1))
	check(err)
	fmt.Println("Int", test)

	grid[0] = append(grid[0], 75)
	grid[1] = append(grid[1], 95, 64)
	grid[2] = append(grid[2], 17, 47, 82)
	grid[3] = append(grid[3], 18, 35, 87, 10)
	grid[4] = append(grid[4], 20, 04, 82, 47, 65)
	grid[5] = append(grid[5], 19, 01, 23, 75, 03, 34)
	grid[6] = append(grid[6], 88, 02, 77, 73, 07, 63, 67)
	grid[7] = append(grid[7], 99, 65, 04, 28, 06, 16, 70, 92)
	grid[8] = append(grid[8], 41, 41, 26, 56, 83, 40, 80, 70, 33)
	grid[9] = append(grid[9], 41, 48, 72, 33, 47, 32, 37, 16, 94, 29)
	grid[10] = append(grid[10], 53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14)
	grid[11] = append(grid[11], 70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57)
	grid[12] = append(grid[12], 91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48)
	grid[13] = append(grid[13], 63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31)
	grid[14] = append(grid[14], 04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23)

	// Make the slice to hold the individual maximum numbers at each index of grid.
	max := make([][]float64, 15)

	// Go through each index starting at the top. Populate the max array with the maximum possible sum at each index by adding the
	// grid number to the maximum of the two maximums for the above indexes. Left and right hand indexes are taken account of.
	for ind := range grid {
		max[ind] = make([]float64, 0)
		for ind2 := range grid[ind] {
			max[ind] = append(max[ind], 0)
			if ind == 0 {
				max[ind][ind2] = grid[ind][ind2]
			} else if ind2 == 0 {
				max[ind][ind2] = grid[ind][ind2] + max[ind-1][0]
			} else if ind == ind2 {
				max[ind][ind2] = grid[ind][ind2] + max[ind-1][ind2-1]
			} else {
				max[ind][ind2] = grid[ind][ind2] + math.Max(max[ind-1][ind2-1], max[ind-1][ind2])
			}
		}
	}

	var maxi float64
	for ind := range max[14] {
		if max[14][ind] > maxi {
			maxi = max[14][ind]
		}
	}
	fmt.Println("Answer", maxi)
}
