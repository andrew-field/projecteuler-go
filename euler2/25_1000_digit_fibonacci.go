/*

The Fibonacci sequence is defined by the recurrence relation:

(x₁²
Fₙ = Fₙ₋₁ + Fₙ₋₂, where F₁ = 1 and F₂ = 1.

Hence the first 12 terms will be:
F₁ = 1
F₂ = 1
F₃ = 2
F₄ = 3
F₅ = 5
F₆ = 8
F₇ = 13
F₈ = 21
F₉ = 34
F₁₀ = 55
F₁₁ = 89
F₁₂ = 144

The 12th term, F₁₂, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

*/

package euler2

import (
	"math/big"

	"github.com/andrew-field/maths/v2"
)

func digitFibonacci(n int) int {
	num1, num2 := big.NewInt(1), big.NewInt(1)
	fib := big.NewInt(1)

	i := 2
	for ; maths.NumberOfDigits(num2) < n; i++ {
		fib.Set(num1)
		num1.Set(num2)
		num2.Add(fib, num2)
	}

	return i
}

// Nothing fancy here, just a simple Fibonacci sequence generator.
// There is some information on the magnitude of Fibonacci numbers: https://en.wikipedia.org/wiki/Fibonacci_sequence#Magnitude
