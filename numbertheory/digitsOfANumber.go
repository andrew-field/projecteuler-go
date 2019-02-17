package numbertheory

import (
	"math/big"
	"strconv"
)

// GetNumberOfDigitsOfAnInt returns the number of digits an int has.
func GetNumberOfDigitsOfAnInt(number int) int {
	if number < 0 {
		number *= -1
	}

	return len(strconv.Itoa(number))
}

// GetNumberOfDigitsOfABigInt returns the number of digits a big.Int has.
func GetNumberOfDigitsOfABigInt(number *big.Int) int {
	// Uses altNumber so as to not change the original number.
	altNumber := big.NewInt(0)
	altNumber.Set(number)
	altNumber.Abs(altNumber)
	return len(altNumber.String())
}

// GetDigitsOfAnInt returns and fills a channel with the digits of a number
// starting with the smallest magnitude numbers (Right to left).
func GetDigitsOfAnInt(number int) chan int {
	if number < 0 {
		number *= -1
	}

	digitsChannel := make(chan int, 10)
	go func() {
		// 456/10 = 45 with int.
		for number > 9 {
			digitsChannel <- number % 10
			number /= 10
		}

		digitsChannel <- number

		close(digitsChannel)
	}()

	return digitsChannel
}

// GetDigitsOfABigInt fills and returns a channel with the digits of a big.Int number
// starting with the smallest magnitude numbers (Right to left).
func GetDigitsOfABigInt(number *big.Int) chan int {
	// Uses altNumber so as to not change the original number.
	altNumber := big.NewInt(0)
	altNumber.Set(number)
	altNumber.Abs(altNumber)

	digitsChannel := make(chan int, 10)
	go func() {
		ten := big.NewInt(10)
		var digit big.Int

		// Dividing these Ints by 10 truncates the decimal places so it is fine.
		for !altNumber.IsInt64() || altNumber.Int64() > 9 {
			// Go is so nice with this handy function that does everything I need.
			// Sets altNumber to altNumber / 10 and sets digit to altNumber mod 10 (The last digit).
			altNumber.QuoRem(altNumber, ten, &digit)
			digitsChannel <- int(digit.Int64())
		}

		digitsChannel <- int(altNumber.Int64())

		close(digitsChannel)
	}()

	return digitsChannel
}

// GetDigitsOfAnIntInASlice returns a slice of the digits of a number as written.
// This function uses string conversions instead of maths.
func GetDigitsOfAnIntInASlice(number int) []int {
	if number < 0 {
		number *= -1
	}

	return digitsOfANumber(strconv.Itoa(number))
}

// GetDigitsOfABigNumberInASlice returns a slice of the digits of a big.Int number as written.
// This function uses string conversions instead of maths.
func GetDigitsOfABigNumberInASlice(number *big.Int) []int {
	// Uses altNumber so as to not change the original number.
	altNumber := big.NewInt(0)
	altNumber.Set(number)
	altNumber.Abs(altNumber)
	return digitsOfANumber(altNumber.Text(10))
}

func digitsOfANumber(number string) []int {
	digits := make([]int, 0)

	for _, val := range number {
		digit, e := strconv.Atoi(string(val))
		if e != nil {
			panic(e)
		}
		digits = append(digits, digit)
	}

	return digits
}
