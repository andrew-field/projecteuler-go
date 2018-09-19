# Testing Go
Testing Go with Project Euler

Library details.

numbertheory:

primeNumbersBelowCeiling:
This program creates a slice of all the numbers from 2 to the max prime size 'ceiling'. It uses a Euclidean sieve/Sieve of Eratosthenes to then find all the primes in the slice.

primeNumbersContinuously:
This uses again an Euclidean sieve/Sieve of Eratosthenes but in parts. The program takes one small slice, the size of which is defined by the user and finds all the primes in that slice before moving on to a new slice. For each new slice it must check all the new numbers against the current prime numbers before sieving.

primeFactorisation:
This program provides a prime factorisation of a number given at runtime. It simply checks whether each prime number in order is a factor and if so, how many times, before moving on to the next prime.

lowestCommonMultiple:
Returns the lowest common multiple of a group of numbers. lcm by raising each prime factor to the maximum number of times that prime factor appears in any of the numbers and then multiplying each of these results.

numberOfDivisors:
Returns the number of divisors a number has using its prime factorisation.

From the prime factorization of n

    n = ∏ i = 1 p i α i ∥ n , α i ≥ 1 ω ( n ) p i α i , {\displaystyle n=\prod _{i=1 \atop {{p_{i}}^{\alpha _{i}}\parallel n,\,\alpha _{i}\geq 1}}^{\omega (n)}{p_{i}}^{\alpha _{i}},\,} {\displaystyle n=\prod _{i=1 \atop {{p_{i}}^{\alpha _{i}}\parallel n,\,\alpha _{i}\geq 1}}^{\omega (n)}{p_{i}}^{\alpha _{i}},\,}

where the pi
are the distinct prime factors of n
and ω (n)
is the number of distinct prime factors of n
, we obtain the number of divisors of n

    σ 0 ( n ) = ∏ i = 1 ω ( n ) ( 1 + α i ) , {\displaystyle \sigma _{0}(n)=\prod _{i=1}^{\omega (n)}(1+\alpha _{i}),\,} {\displaystyle \sigma _{0}(n)=\prod _{i=1}^{\omega (n)}(1+\alpha _{i}),\,}

since for each pi
we can choose the exponent from 0
to αi
to build a divisor of n
.
