# Testing Go

## Testing Go with Project Euler.  
This is a repository to keep my code for solving project Euler challenges.

## Why is this here?  
1. I would like to learn Golang and solve these challenges.
2.
![Because sharing (after solving for yourself) is caring.](images/reason2.jpg)


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
See [this brief explantion](https://oeis.org/wiki/Number_of_divisors_function "Wiki article")
