package projecteuler

// MultiplesOf3And5MethodOne sums the multiples of 3 and 5 below 1000 without unnecessary calculations.
// The question is whether the steps taken to avoid frivolous calculations actually make the process quicker
// in the end and if they do, is the effort worth the benefit.
func MultiplesOf3And5MethodOne() int {
	// Total sum.
	total := 0

	// Include multiple of 3.
	for i := 3; i < 1000; i += 3 {
		total += i
	}

	// Include multiples of 5.
	i := 1
	for j := 5; j < 1000; j += 5 {
		// Exclude multiples of 3.
		if i == 3 {
			i = 1
			continue
		}
		total += j
		i++
	}

	return total
}

// MultiplesOf3And5MethodTwo sums the multiples of 3 and 5 below 1000 in the simplist way, including
// wasteful calculations.
func MultiplesOf3And5MethodTwo() int {
	// Total sum.
	total := 0

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	return total
}
