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
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 55},
	}
	firstOldest := slices.MaxFunc(people, func(a, b Person) int {
		return cmp.Compare(a.Age, b.Age)
	})
	fmt.Println(firstOldest.Name)
}
