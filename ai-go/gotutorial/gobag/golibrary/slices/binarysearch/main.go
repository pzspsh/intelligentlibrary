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
	names := []string{"Alice", "Bob", "Vera"}
	n, found := slices.BinarySearch(names, "Vera")
	fmt.Println("Vera:", n, found)
	n, found = slices.BinarySearch(names, "Bill")
	fmt.Println("Bill:", n, found)
}
