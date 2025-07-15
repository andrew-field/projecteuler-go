/*
Using the given names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

package euler2

import (
	"encoding/csv"
	"os"
	"slices"
)

// nameScores returns the summation of all the name scores in the file p022_names.txt
func nameScores() int {
	letterValues := map[rune]int{
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
	slices.Sort(names) // Score depends on position in list.

	// Calculate scores.
	total := 0
	for ind, name := range names {
		nameScore := 0
		for _, letter := range name {
			nameScore += letterValues[letter]
		}
		total += nameScore * (ind + 1)
	}

	return total
}

// Could have a goroutine for each name to calculate the score concurrently, but this is fine.
