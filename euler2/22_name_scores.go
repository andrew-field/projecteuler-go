package euler2

import (
	"encoding/csv"
	"os"
	"sort"
	"sync"
)

var (
	mu           = sync.Mutex{}
	letterValues = map[rune]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'D': 4,
		'E': 5,
		'F': 6,
		'G': 7,
		'H': 8,
		'I': 9,
		'J': 10,
		'K': 11,
		'L': 12,
		'M': 13,
		'N': 14,
		'O': 15,
		'P': 16,
		'Q': 17,
		'R': 18,
		'S': 19,
		'T': 20,
		'U': 21,
		'V': 22,
		'W': 23,
		'X': 24,
		'Y': 25,
		'Z': 26,
	}
)

// nameScores returns the summation of all the name scores in the file p022_names.txt
func nameScores() int {
	f, err := os.Open("p022_names.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read names and sort.
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	names := records[0]
	sort.Strings(names) // Score depends on position in list.

	// Send scores.
	scores := make(chan int)
	for ind, val := range names {
		go func() { // Calculate and send scores.
			nameScore := 0
			for _, letter := range val {
				mu.Lock()
				nameScore += letterValues[letter]
				mu.Unlock()
			}
			scores <- nameScore * (ind + 1)
		}()
	}

	// Sum scores as received.
	total := 0
	for range names { // There should be len(names) scores so no need for wait groups.
		total += <-scores
	}
	return total
}

// It would probably be fine without the concurrency.
