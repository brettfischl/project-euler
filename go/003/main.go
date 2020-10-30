package main

import (
	"fmt"
)

// https://projecteuler.net/problem=3
func main() {
	fmt.Println("What is the largest prime factor of the number 600851475143?")

	num := 600851475143

	var primeFactors []int

	for num%2 == 0 {
		num = num / 2
		primeFactors = append(primeFactors, 2)
	}

	primeFactorization(3, num, &primeFactors)

	fmt.Println(primeFactors)
}

func primeFactorization(n int, target int, factors *[]int) {
	if target%n == 0 {
		*factors = append(*factors, n)
		target = target / n
	}

	if n*n > target {
		if n > 2 {
			*factors = append(*factors, target)
		}
		return
	}

	primeFactorization(n+2, target, factors)
}
