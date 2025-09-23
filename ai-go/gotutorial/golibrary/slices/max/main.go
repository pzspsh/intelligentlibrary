package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, -10, 8}
	fmt.Println(slices.Max(numbers))
}
