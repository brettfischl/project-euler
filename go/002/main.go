package main

import "fmt"

// https://projecteuler.net/problem=2
func main() {
	fmt.Println("By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.")

	n1 := 1
	n2 := 2

	end := 4000000
	sum := 0
	for n1 < end {
		if n1%2 == 0 {
			sum += n1
		}
		tmp := n2
		n2 = n1 + n2
		n1 = tmp
	}

	fmt.Println(sum)
}

func fib(i int, n int) {

}
