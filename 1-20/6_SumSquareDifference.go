package euler1

// SumSquareDifference returns the difference between the sum of the squares of xₖ, xₖ₊₁, ..., xₙ and the square of the sum, where xₖ₊₁ = xₖ + 1 and xₖ = x.
// Assumes x >= 0 and n >= 0
func SumSquareDifference(x, n int) int {
	return n*x*(x+n-1)*(n-1) + (n*(3*(n*n*n)-10*(n*n)+(9*n)-2))/12
}

// 1+2+...+n = n(n+1)/2
// 1²+2²+...+n² = n(n+1)(2n+1)/6

// So for xₖ₊₁ = xₖ + 1,
// (x₁ + ... +xₙ)² = n²x₁² + n²(n-1)x₁ + (n²(n-1)²)/4
// x₁² + ... + xₙ² = nx₁² + n(n-1)x₁ + (n(n-1)(2n-1))/6
// (x₁ + ... +xₙ)² - (x₁² + ... + xₙ²) = nx₁(x₁+n-1)(n-1) + (n(3n³-10n²+9n-2))/12
