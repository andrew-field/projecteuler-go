package numbertheory

import (
	"math"
)

// GetPrimeNumbersContinuously fills a channel with the prime numbers.
// The smaller the slice increment size, the quicker the primes will come initially but the slower the primes will come over time.
// The larger the slice increment size, the slower the primes will come initially but the more suitable it will be for larger primes.
// Things will probably break at the extremities with accuracy issues surrounding float64/uint.
// Syncing and safely exiting this function can be done through sending a signal on the done channel.
func GetPrimeNumbersContinuously(sliceIncrementsSize int) (chan uint, chan bool) {
	if sliceIncrementsSize < 1 {
		panic("The slice increment size must be larger than 0.")
	}

	primeChannel := make(chan uint, 100)
	doneChannel := make(chan bool)

	go func() {
		defer func() {
			close(primeChannel)
		}()

		// Inital slice for primes.
		primes := make([]uint, 0)

		// Setup.
		var firstNumberToCheck uint = 2

		// Slice containing the new numbers to check for primality.
		newNumbers := make([]uint, sliceIncrementsSize)

		for {
			// Make new numbers.
			for ind := range newNumbers {
				newNumbers[ind] = firstNumberToCheck
				firstNumberToCheck++

				// Check for finish (overflow).
				if firstNumberToCheck < 0 {
					newNumbers = newNumbers[:ind]
					generatePrimes(&primes, &newNumbers, primeChannel, doneChannel, ind)
					return
				}
			}

			// Generate primes. This function will return true if a signal on the done channel comes through.
			if generatePrimes(&primes, &newNumbers, primeChannel, doneChannel, sliceIncrementsSize) {
				return
			}
		}
	}()

	return primeChannel, doneChannel
}

// Basic euclidean sieve that probably has some issues somewhere with the type casting...
func generatePrimes(primes *[]uint, newNumbers *[]uint, primeChannel chan uint, doneChannel chan bool, sliceIncrementsSize int) bool {
	for ind, val := range *newNumbers {
		if val != 1 {
			isPrime := true
			for _, primeValue := range *primes {
				if float64(primeValue) > math.Floor(math.Sqrt(float64(val))) {
					break
				}
				if val%primeValue == 0 {
					for ind < sliceIncrementsSize {
						if (*newNumbers)[ind] != 1 {
							(*newNumbers)[ind] = 1
						}
						ind += int(primeValue)
					}
					isPrime = false
					break
				}
			}

			if isPrime {
				select {
				// If finished then end the function.
				case <-doneChannel:
					return true
				case primeChannel <- val:
					*primes = append(*primes, val)
					break
				}
			}
		}
	}

	return false
}
