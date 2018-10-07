package main

import "fmt"

const one, two, three, four, five, six, seven, eight, nine = 3, 3, 5, 4, 4, 3, 5, 5, 4
const ten, eleven, twelve, thirteen, fourteen, fifteen, sixteen, seventeen, eighteen, nineteen = 3, 6, 6, 8, 8, 7, 7, 9, 8, 8
const twenty, thirty, forty, fifty, sixty, seventy, eighty, ninety = 6, 6, 5, 5, 5, 7, 6, 6
const onehundred, twohundred, threehundred, fourhundred, fivehundred = 10, 10, 12, 11, 11
const sixhundred, sevenhundred, eighthundred, ninehundred = 10, 12, 12, 11
const and = 3
const onethousand = 11

func main() {

	// Total.
	total := 0

	// Each of these digits appears 90 times.
	total += 90 * (one + two + three + four + five + six + seven + eight + nine)

	// Each of these digits appears 10 times.
	total += 10 * (ten + eleven + twelve + thirteen + fourteen + fifteen + sixteen + seventeen + eighteen + nineteen)

	// Each of these digits appears 100 times.
	total += 100 * (twenty + thirty + forty + fifty + sixty + seventy + eighty + ninety)
	total += 100 * (onehundred + twohundred + threehundred + fourhundred + fivehundred + sixhundred + sevenhundred + eighthundred + ninehundred)

	// The word and appears 891 times.
	total += 891 * and

	total += onethousand

	fmt.Println("total:", total)
}
