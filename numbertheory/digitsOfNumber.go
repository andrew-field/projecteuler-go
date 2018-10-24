package numbertheory

import (
	"math/big"
	"strconv"
)

// GetDigitsOfABigNumber fills and returns a channel with the digits of a big.Int number
// starting with the smallest magnitude numbers (Right to left).
func GetDigitsOfABigNumber(number big.Int) chan int {

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

// GetDigitsOfANumber fills and returns a channel with the digits of a number
// starting with the smallest magnitude numbers (Right to left).
func GetDigitsOfANumber(number int) chan int {

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

// GetDigitsOfABigNumberInSlice returns a slice of the digits of a big.Int number as written.
func GetDigitsOfABigNumberInSlice(number big.Int) []int {
	if number.Sign() == -1 {
		panic("The number should be 0 or positive.")
	}

	return digitsOfANumber(number.Text(10))
}

// GetDigitsOfANumberInSlice returns a slice of the digits of a number as written.
func GetDigitsOfANumberInSlice(number int) []int {
	if number < 0 {
		panic("The number should be 0 or positive.")
	}

	return digitsOfANumber(strconv.Itoa(number))
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
