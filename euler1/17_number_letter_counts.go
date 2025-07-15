/*
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
*/

package euler1

const one, two, three, four, five, six, seven, eight, nine = 3, 3, 5, 4, 4, 3, 5, 5, 4
const ten, eleven, twelve, thirteen, fourteen, fifteen, sixteen, seventeen, eighteen, nineteen = 3, 6, 6, 8, 8, 7, 7, 9, 8, 8
const twenty, thirty, forty, fifty, sixty, seventy, eighty, ninety = 6, 6, 5, 5, 5, 7, 6, 6
const onehundred, twohundred, threehundred, fourhundred, fivehundred = 10, 10, 12, 11, 11
const sixhundred, sevenhundred, eighthundred, ninehundred = 10, 12, 12, 11
const and = 3
const onethousand = 11

// numberLetterCounts returns how many letters would be used if all the numbers from 1 to 1000 (one thousand) inclusive were written out in words.
func numberLetterCounts() int {
	// Each of these digits appears 90 times.
	total := 90 * (one + two + three + four + five + six + seven + eight + nine)

	// Each of these digits appears 10 times.
	total += 10 * (ten + eleven + twelve + thirteen + fourteen + fifteen + sixteen + seventeen + eighteen + nineteen)

	// Each of these digits appears 100 times.
	total += 100 * (twenty + thirty + forty + fifty + sixty + seventy + eighty + ninety)
	total += 100 * (onehundred + twohundred + threehundred + fourhundred + fivehundred + sixhundred + sevenhundred + eighthundred + ninehundred)

	// The word "and" appears 891 times.
	total += 891 * and

	return total + onethousand
}

// Simple brute force. Each number has as value and appears a certain number of times. Not sure if there is a clever way of doing it.
