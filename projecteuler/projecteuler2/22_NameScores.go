package projecteuler2

import (
	"bufio"
	"encoding/csv"
	"os"
	"path/filepath"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// NameScores returns the total of all the name scores in the file p022_names.txt
func NameScores() int {
	// Open file.
	absPath, err := filepath.Abs("p022_names.txt")
	check(err)
	file, err := os.Open(absPath)
	check(err)
	defer func() {
		file.Close()
	}()
	reader := csv.NewReader(bufio.NewReader(file))
	records, err := reader.ReadAll()
	check(err)
	names := records[0]
	sort.Strings(names)
	letterValues := map[string]int{
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

	total := 0
	for ind, val := range names {
		nameScore := 0
		for _, val := range val {
			nameScore += letterValues[string(val)]
		}
		total += (nameScore * (ind + 1))
	}

	return total
}
