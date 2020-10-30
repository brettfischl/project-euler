package main

import "fmt"

// https://projecteuler.net/problem=1
func main() {
	fmt.Println("Find the sum of all the multiples of 3 or 5 below 1000.")

	i := 999
	sum := 0
	for i > 0 {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
		i--
	}

	fmt.Println(sum)
}
