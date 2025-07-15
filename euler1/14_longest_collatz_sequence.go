/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package euler1

// Contains the critical information about a Collatz sequence.
type Sequence struct {
	startingNumber, length int
}

// longestCollatzSequence returns the starting number, under one million, which produces the longest chain/Collatz sequence.
func longestCollatzSequence() int {
	// Keep track of the longest sequence.
	longestSequence := Sequence{}

	// Channel to send the sequence information once it has been calculated.
	sequenceCh := make(chan Sequence)

	// Generate all of the sequences concurrently.
	for startingNumber := 500000; startingNumber < 1000000; startingNumber++ {
		go generateSequence(sequenceCh, startingNumber)
	}

	for range 499999 { // 499999 sequences to check.
		seqResult := <-sequenceCh
		if seqResult.length > longestSequence.length {
			longestSequence = seqResult
		}
	}
	return longestSequence.startingNumber
}

func generateSequence(ch chan<- Sequence, num int) {
	seq := Sequence{num, 1} // Set the starting number and the length of the sequence to 1 but will reuse the num variable.
	for ; num != 1; seq.length++ {
		if num%2 == 0 {
			num /= 2
		} else {
			num = 3*num + 1 // The next term must be even so it is possible to save some time here.
			seq.length++
			num /= 2
		}
	}
	ch <- seq
}

// My previous implementation kept track of how many sequences can be skipped because the starting number appeared in a previous sequence and hence the
// new sequence would be necessarily shorter. This was refactored to have a concurrent design and to make it simple/easy to read, though it no longer has this fine tuning
// and calculates every sequence. Still, the first half of all the sequences starting under 1'000'000 will appear as part of a longer sequence starting in the second half,
// so the first 50% can be skipped.
