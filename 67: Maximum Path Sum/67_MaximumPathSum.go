package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Height of the pyramid/Length of longest row.
	length := 100

	// Make the grid for the numbers.
	pyramid := make([][][]float64, length)
	// Each index of the pyramid has a corresponding "max slot".
	for ind := range pyramid {
		pyramid[ind] = make([][]float64, 0)
	}

	// Open file.
	absPath, err := filepath.Abs("p067_triangle.txt")
	check(err)
	f, err := os.Open(absPath)
	check(err)
	defer func() {
		f.Close()
	}()
	reader := bufio.NewReader(f)

	// Read 2 digits every time and populate the value slots.
	number := make([]byte, 2)
	for i := 0; i < length; i++ {
		for j := 0; j <= i; j++ {
			_, err = reader.Read(number)
			check(err)
			valueSlot, err := strconv.ParseFloat(string(number), 64)
			check(err)
			pyramid[i] = append(pyramid[i], []float64{valueSlot, 0})
		}
	}

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
