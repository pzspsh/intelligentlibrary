package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(slices.IsSorted([]string{"Alice", "Bob", "Vera"}))
	fmt.Println(slices.IsSorted([]int{0, 2, 1}))
}
