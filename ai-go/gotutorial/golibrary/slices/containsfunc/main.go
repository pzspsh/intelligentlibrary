/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 22:03:25
*/
package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, -10, 8}
	hasNegative := slices.ContainsFunc(numbers, func(n int) bool {
		return n < 0
	})
	fmt.Println("Has a negative:", hasNegative)
	hasOdd := slices.ContainsFunc(numbers, func(n int) bool {
		return n%2 != 0
	})
	fmt.Println("Has an odd number:", hasOdd)
}
