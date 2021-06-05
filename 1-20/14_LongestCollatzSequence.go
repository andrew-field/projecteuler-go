package euler1

type sequence struct {
	startingNumber, length int
}

// LongestCollatzSequence returns the starting number, under one million, which produces the longest chain/Collatz sequence.
func LongestCollatzSequence() int {

	longest := sequence{}
	sequenceCh := make(chan sequence)

	var seqResult sequence
	doneCh := make(chan bool)
	go func() {
		for i := 1; i < 500000; i++ {
			seqResult = <-sequenceCh
			if seqResult.length > longest.length {
				longest = seqResult
			}
		}
		doneCh <- true
	}()

	for i := 500000; i < 1000000; i++ {
		go generateSequence(sequenceCh, i)
	}

	<-doneCh
	return longest.startingNumber
}

func generateSequence(ch chan sequence, num int) {
	seq := sequence{num, 1}
	for ; num != 1; seq.length++ {
		if num%2 == 0 {
			num /= 2
		} else {
			num = 3*num + 1 // The next term must be even so it is possible to save some time here but omitted to keep things simple.
		}
	}
	ch <- seq
}

// My previous implementation kept track of how many sequences can be skipped because the starting number appeared in a previous sequence and hence the
// new sequence would be necessarily shorter. This was refactored to have a concurrent design and to make it simple/easy to read, though it no longer has this fine tuning
// and calculates every sequence. Still, the first half of all the sequences starting under 1'000'000 will appear as part of a longer sequence starting in the second half, so the first 50% can be skipped.
