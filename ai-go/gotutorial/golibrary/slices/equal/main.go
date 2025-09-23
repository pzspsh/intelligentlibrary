package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, 8}
	fmt.Println(slices.Equal(numbers, []int{0, 42, 8}))
	fmt.Println(slices.Equal(numbers, []int{10}))
}
