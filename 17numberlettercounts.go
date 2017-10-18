package main

import "fmt"

const one, two, three, four, five, six, seven, eight, nine = 3, 3, 5, 4, 4, 3, 5, 5, 4
const ten, eleven, twelve, thirteen, fourteen, fifteen, sixteen, seventeen, eighteen, nineteen = 3, 6, 6, 8, 8, 7, 7, 9, 8, 8
const twenty, thirty, forty, fifty, sixty, seventy, eighty, ninety = 6, 6, 5, 5, 5, 7, 6, 6
const onehundredand, twohundredand, threehundredand, fourhundredand, fivehundredand = 13, 13, 15, 14, 14
const sixhundredand, sevenhundredand, eighthundredand, ninehundredand = 13, 15, 15, 14
const onethousand = 11

func main() {
	total := 0

	total += 90 * one
	total += 90 * two
	total += 90 * three
	total += 90 * four
	total += 90 * five
	total += 90 * six
	total += 90 * seven
	total += 90 * eight
	total += 90 * nine

	total += 10 * ten
	total += 10 * eleven
	total += 10 * twelve
	total += 10 * thirteen
	total += 10 * fourteen
	total += 10 * fifteen
	total += 10 * sixteen
	total += 10 * seventeen
	total += 10 * eighteen
	total += 10 * nineteen

	total += 100 * twenty
	total += 100 * thirty
	total += 100 * forty
	total += 100 * fifty
	total += 100 * sixty
	total += 100 * seventy
	total += 100 * eighty
	total += 100 * ninety

	total += 100 * onehundredand
	total += 100 * twohundredand
	total += 100 * threehundredand
	total += 100 * fourhundredand
	total += 100 * fivehundredand
	total += 100 * sixhundredand
	total += 100 * sevenhundredand
	total += 100 * eighthundredand
	total += 100 * ninehundredand

	// Too many "and".
	total -= 9 * 3

	total += onethousand

	fmt.Println("total:", total)
}
