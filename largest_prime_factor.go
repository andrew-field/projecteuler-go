package main

import "fmt"

func main () {

num := 600851475143

primes := make([]int,10000)

for i:=0;i<10000;i++ {
    primes[i] = i+1
}

k := 0

for j:=2;j<=100;j++ {
    for k=1;k<10000;k++ {
        if primes[k] != 1 && primes[k]%j == 0 && primes[k]/j != 1 {
            primes[k] = 1
        }
    }
}

factors := make([]int,0)

for i:=0;i<10000;i++ {
    if primes[i]!=1 && num%primes[i] == 0 {
        factors = append(factors, primes[i])
        num /= primes[i]
        i--
        if num==1 {
            break        
        }
    }
}

fmt.Println(factors)

length := len(factors)
largest := factors[0]

for i:=1;i<length;i++ {
    if factors[i] > largest {
        largest = factors[i]
    }
}

fmt.Println(largest)

}
