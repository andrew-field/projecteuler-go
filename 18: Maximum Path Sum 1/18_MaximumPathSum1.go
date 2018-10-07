package main

import (
	"fmt"
	"math"
)

func main() {

	// Height of pyramid/Length of longest row.
	length := 15

	// Make the grid for the numbers.
	pyramid := make([][][]float64, length)

	// Each index of the pyramid has a corresponding "max slot" (The zeroes).
	pyramid[0] = [][]float64{{75, 0}}
	pyramid[1] = [][]float64{{95, 0}, {64, 0}}
	pyramid[2] = [][]float64{{17, 0}, {47, 0}, {82, 0}}
	pyramid[3] = [][]float64{{18, 0}, {35, 0}, {87, 0}, {10, 0}}
	pyramid[4] = [][]float64{{20, 0}, {04, 0}, {82, 0}, {47, 0}, {65, 0}}
	pyramid[5] = [][]float64{{19, 0}, {01, 0}, {23, 0}, {75, 0}, {03, 0}, {34, 0}}
	pyramid[6] = [][]float64{{88, 0}, {02, 0}, {77, 0}, {73, 0}, {07, 0}, {63, 0}, {67, 0}}
	pyramid[7] = [][]float64{{99, 0}, {65, 0}, {04, 0}, {28, 0}, {06, 0}, {16, 0}, {70, 0}, {92, 0}}
	pyramid[8] = [][]float64{{41, 0}, {41, 0}, {26, 0}, {56, 0}, {83, 0}, {40, 0}, {80, 0}, {70, 0}, {33, 0}}
	pyramid[9] = [][]float64{{41, 0}, {48, 0}, {72, 0}, {33, 0}, {47, 0}, {32, 0}, {37, 0}, {16, 0}, {94, 0}, {29, 0}}
	pyramid[10] = [][]float64{{53, 0}, {71, 0}, {44, 0}, {65, 0}, {25, 0}, {43, 0}, {91, 0}, {52, 0}, {97, 0}, {51, 0}, {14, 0}}
	pyramid[11] = [][]float64{{70, 0}, {11, 0}, {33, 0}, {28, 0}, {77, 0}, {73, 0}, {17, 0}, {78, 0}, {39, 0}, {68, 0}, {17, 0}, {57, 0}}
	pyramid[12] = [][]float64{{91, 0}, {71, 0}, {52, 0}, {38, 0}, {17, 0}, {14, 0}, {91, 0}, {43, 0}, {58, 0}, {50, 0}, {27, 0}, {29, 0}, {48, 0}}
	pyramid[13] = [][]float64{{63, 0}, {66, 0}, {04, 0}, {68, 0}, {89, 0}, {53, 0}, {67, 0}, {30, 0}, {73, 0}, {16, 0}, {69, 0}, {87, 0}, {40, 0}, {31, 0}}
	pyramid[14] = [][]float64{{04, 0}, {62, 0}, {98, 0}, {27, 0}, {23, 0}, {9, 0}, {70, 0}, {98, 0}, {73, 0}, {93, 0}, {38, 0}, {53, 0}, {60, 0}, {04, 0}, {23, 0}}

	// Go through each index starting at the top.
	// Populate each max slot of each index with the maximum possible sum with which it is possible
	// to reach that index. Do this by adding the grid number at the index to the maximum of the
	// two max slots for the above indexes (Directly above or left).
	for ind := range pyramid {
		// The top.
		if ind == 0 {
			pyramid[ind][0][1] = pyramid[ind][0][0]
			continue
		}
		for ind2 := range pyramid[ind] {
			if ind2 == 0 { // The leftmost indexes.
				pyramid[ind][ind2][1] = pyramid[ind][ind2][0] + pyramid[ind-1][ind2][1]
			} else if ind == ind2 { // The rightmost indexes.
				pyramid[ind][ind2][1] = pyramid[ind][ind2][0] + pyramid[ind-1][ind2-1][1]
			} else { // The rest.
				pyramid[ind][ind2][1] = pyramid[ind][ind2][0] + math.Max(pyramid[ind-1][ind2-1][1], pyramid[ind-1][ind2][1])
			}
		}
	}

	// Maximum.
	var max float64

	// Find the maximum of the max slots in the final row.
	for _, val := range pyramid[length-1] {
		max = math.Max(max, val[1])
	}

	fmt.Println("Answer:", max)
}
