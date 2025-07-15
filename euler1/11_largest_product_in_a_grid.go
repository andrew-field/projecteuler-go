/*
In the 20×20 grid below, four numbers along a diagonal line have been marked in red.

08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00
81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65
52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91
22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80
24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50
32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70
67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21
24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72
21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95
78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92
16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57
86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58
19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40
04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66
88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69
04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36
20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16
20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54
01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48

The product of these numbers is 26 × 63 × 78 × 14 = 1788696.

What is the greatest product of four adjacent numbers in the same direction (up, down, left, right, or diagonally) in the 20×20 grid?
*/

package euler1

import (
	"sync"
)

// largestProductInAGrid returns the greatest product of four adjacent numbers
// in the same direction (up, down, left, right, or diagonally) in a grid.
func largestProductInAGrid() int {
	matrix := make([][]int, 20)

	// Setup.
	matrix[0] = []int{8, 2, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8}
	matrix[1] = []int{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 04, 56, 62, 0}
	matrix[2] = []int{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 03, 49, 13, 36, 65}
	matrix[3] = []int{52, 70, 95, 23, 04, 60, 11, 42, 69, 24, 68, 56, 01, 32, 56, 71, 37, 02, 36, 91}
	matrix[4] = []int{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80}
	matrix[5] = []int{24, 47, 32, 60, 99, 03, 45, 02, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50}
	matrix[6] = []int{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70}
	matrix[7] = []int{67, 26, 20, 68, 02, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21}
	matrix[8] = []int{24, 55, 58, 05, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72}
	matrix[9] = []int{21, 36, 23, 9, 75, 0, 76, 44, 20, 45, 35, 14, 0, 61, 33, 97, 34, 31, 33, 95}
	matrix[10] = []int{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 03, 80, 04, 62, 16, 14, 9, 53, 56, 92}
	matrix[11] = []int{16, 39, 05, 42, 96, 35, 31, 47, 55, 58, 88, 24, 0, 17, 54, 24, 36, 29, 85, 57}
	matrix[12] = []int{86, 56, 0, 48, 35, 71, 89, 07, 05, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58}
	matrix[13] = []int{19, 80, 81, 68, 05, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 04, 89, 55, 40}
	matrix[14] = []int{04, 52, 8, 83, 97, 35, 99, 16, 07, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66}
	matrix[15] = []int{88, 36, 68, 87, 57, 62, 20, 72, 03, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69}
	matrix[16] = []int{04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36}
	matrix[17] = []int{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 04, 36, 16}
	matrix[18] = []int{20, 73, 35, 29, 78, 31, 90, 01, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 05, 54}
	matrix[19] = []int{01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 01, 89, 19, 67, 48}

	// Maximum.
	var max0 int
	var max1 int
	var max2 int
	var max3 int

	var wg sync.WaitGroup
	wg.Add(4)

	// Four calculations are wasted for every zero but it is simple enough.
	go func() {
		// Horizontal.
		for j := range 20 {
			for i := range 17 {
				max0 = max(max0, matrix[j][i]*matrix[j][i+1]*matrix[j][i+2]*matrix[j][i+3]) // Can use builtin max function.
			}
		}
		wg.Done()
	}()

	go func() {
		// Vertical.
		for j := range 17 {
			for i := range 20 {
				max1 = max(max1, matrix[j][i]*matrix[j+1][i]*matrix[j+2][i]*matrix[j+3][i])
			}
		}
		wg.Done()
	}()

	go func() {
		// Left diagonal.
		for j := range 17 {
			for i := range 17 {
				max2 = max(max2, matrix[j][i]*matrix[j+1][i+1]*matrix[j+2][i+2]*matrix[j+3][i+3])
			}
		}
		wg.Done()
	}()

	go func() {
		// Right diagonal.
		for j := range 17 {
			for i := range 17 {
				max3 = max(max3, matrix[j][19-i]*matrix[j+1][18-i]*matrix[j+2][17-i]*matrix[j+3][16-i])
			}
		}
		wg.Done()
	}()
	wg.Wait()

	return max(max0, max1, max2, max3)
}

// This is a simple brute force but at least with some concurrency.
// Only one matrix is needed as there can not be a race condition involving only read operations.
// One could probably use a single max variable but actually then most of the calculations would be locked and the benefit of concurrency negated. This way is simpler.
// All the different products could be stored in one slice and then passed to the max function so the max function would only be called once, but this is simple enough.
