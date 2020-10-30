package main

import (
	"fmt"
	"math"
)

// https://projecteuler.net/problem=4
func main() {
	fmt.Println("Find the largest palindrome made from the product of two 3-digit numbers.")
	n1 := 999
	n2 := 999

	product := findHighestPalindrome(n1, n2)

	fmt.Println(product)
}

func isPalindrome(n int) bool {
	nStr := fmt.Sprint(n)

	result := true
	strLen := len(nStr) - 1
	midPoint := int(math.Round(float64(strLen)/2.0)) - 1
	for i := range nStr {
		if i > midPoint {
			break
		}
		if string(nStr[i]) != string(nStr[strLen-i]) {
			result = false
			break
		}
	}

	return result
}

func findHighestPalindrome(startN1, startN2 int) int {
	n1, n2 := startN1, startN2

	highest := 0
	product := 0
	for n1 >= 100 {
		for n2 >= 100 {
			product = n1 * n2

			if isPalindrome(product) && product > highest {
				highest = product
			}
			n2--
		}
		n2 = startN2
		n1--
	}

	return highest
}
