package numbertheory

import (
	"math"
)

var primes []uint
var newNumbers []uint
var firstNumberToCheck uint
var sizeOfSlice int
var isPrime bool

// GetPrimeNumbersContinuously fills a channel with the prime numbers. This process will buffer through the channel passed in.
// The smaller the slice increment size, the quicker the primes will come initially but the slower the primes will come over time.
// The larger the slice increment size, the slower the primes will come initially but the more suitable it will be for larger primes.
// Things will probably break at the extremities with accuracy issues surrounding float64/uint.
func GetPrimeNumbersContinuously(primeChannel chan uint, sliceIncrementsSize int, doneChannel chan bool) {

	if sliceIncrementsSize < 1 {
		panic("The slice increment size must be larger than 0.")
	}

	// Inital slice for primes.
	primes = make([]uint, 0)
	firstNumberToCheck = 2

	// Slice containing the new numbers to check for primality.
	newNumbers = make([]uint, sliceIncrementsSize)

	for {
		select {
		case <-doneChannel:
			close(primeChannel)
			return
		default:
			// Make new numbers.
			for ind := range newNumbers {
				newNumbers[ind] = firstNumberToCheck
				firstNumberToCheck++

				// Check for finish.
				if firstNumberToCheck < 0 {
					newNumbers = newNumbers[:ind]
					generatePrimes(primeChannel, ind)
					close(primeChannel)
					return
				}
			}

			// Generate primes.
			generatePrimes(primeChannel, sliceIncrementsSize)
		}
	}
}

// Basic euclidean sieve that probably has some issues somewhere with the type casting on line 46...
func generatePrimes(primeChannel chan uint, sliceSize int) {
	for ind, val := range newNumbers {
		isPrime = true
		if val != 1 {
			for _, primeValue := range primes {
				if float64(primeValue) > math.Floor(math.Sqrt(float64(val))) {
					break
				}
				if val%primeValue == 0 {
					for ind < sliceSize {
						if newNumbers[ind] != 1 {
							newNumbers[ind] = 1
						}
						ind += int(primeValue)
					}
					isPrime = false
					break
				}
			}

			if isPrime {
				primeChannel <- val
				primes = append(primes, val)
			}
		}
	}
}
