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
	absNumber := *number
	return len(absNumber.Abs(&absNumber).String())
}

// GetDigitsOfAnInt fills and returns a channel with the digits of a number
// starting with the smallest magnitude numbers (Right to left).
// Syncing and safely exiting this function can be done through flushing the digits channel.
func GetDigitsOfAnInt(number int) chan int {

	if number < 0 {
		panic("The number should be 0 or positive.")
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
// Syncing and safely exiting this function can be done through flushing the digits channel.
func GetDigitsOfABigInt(number big.Int) chan int {

	if number.Sign() == -1 {
		panic("The number should be 0 or positive.")
	}

	digitsChannel := make(chan int, 10)
	go func() {
		ten := big.NewInt(10)
		var digit big.Int

		// Dividing these Ints by 10 truncates the decimal places so it is fine.
		for !number.IsInt64() || number.Int64() > 9 {
			// Go is so nice with this handy function that does everything I need.
			// Sets number to number / 10 and sets digit to number mod 10 (The last digit).
			number.QuoRem(&number, ten, &digit)
			digitsChannel <- int(digit.Int64())
		}

		digitsChannel <- int(number.Int64())

		close(digitsChannel)
	}()

	return digitsChannel
}

// GetDigitsOfAnIntInSlice returns a slice of the digits of a number as written.
func GetDigitsOfAnIntInSlice(number int) []int {
	if number < 0 {
		panic("The number should be 0 or positive.")
	}

	return digitsOfANumber(strconv.Itoa(number))
}

// GetDigitsOfABigNumberInSlice returns a slice of the digits of a big.Int number as written.
func GetDigitsOfABigNumberInSlice(number big.Int) []int {
	if number.Sign() == -1 {
		panic("The number should be 0 or positive.")
	}

	return digitsOfANumber(number.Text(10))
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
