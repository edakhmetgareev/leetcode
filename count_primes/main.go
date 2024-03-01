package main

import (
	"fmt"
)

func main() {
	r := countPrimes(499979)
	fmt.Printf("result: %+v", r)

}

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	primes := make([]bool, n)
	for i := 2; i < n; i++ {
		primes[i] = true
	}

	for i := 0; i*i < n; i++ {
		if primes[i] {
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}

	cnt := 0
	for i := range primes {
		if primes[i] {
			cnt++
		}
	}

	return cnt
}
