package euler2

import (
	"encoding/csv"
	"sort"
	"sync"

	"github.com/andrew-field/projecteuler-go/filefunctions"
)

var (
	mu           = sync.Mutex{}
	letterValues = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
		"J": 10,
		"K": 11,
		"L": 12,
		"M": 13,
		"N": 14,
		"O": 15,
		"P": 16,
		"Q": 17,
		"R": 18,
		"S": 19,
		"T": 20,
		"U": 21,
		"V": 22,
		"W": 23,
		"X": 24,
		"Y": 25,
		"Z": 26,
	}
)

// nameScores returns the summation of all the name scores in the file p022_names.txt
func nameScores() int {
	f := filefunctions.OpenFile("p022_names.txt")
	defer filefunctions.CloseFile(f)

	// Read names and sort.
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	filefunctions.Check(err)
	names := records[0]
	sort.Strings(names) // Score depends on position in list.

	// Sum scores as received.
	total := 0
	scores := make(chan int)
	doneCh := make(chan bool)
	go func() {
		for i := 0; i < len(names); i++ {
			total += <-scores
		}
		doneCh <- true
	}()

	// Send scores.
	for ind, val := range names {
		go func(index int, name string) { // Calculate and send scores
			nameScore := 0
			for _, val := range name {
				mu.Lock()
				nameScore += letterValues[string(val)]
				mu.Unlock()
			}
			scores <- nameScore * (index + 1)
		}(ind, val)
	}

	<-doneCh
	return total
}

// It would probably be fine without the concurrency.
