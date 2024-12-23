/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 22:03:25
*/
package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Gopher", 13},
		{"Alice", 20},
		{"Bob", 24},
		{"Alice", 55},
	}
	// Stable sort by name, keeping age ordering of Alices intact
	slices.SortStableFunc(people, func(a, b Person) int {
		return cmp.Compare(a.Name, b.Name)
	})
	fmt.Println(people)
}
