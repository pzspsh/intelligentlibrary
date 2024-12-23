/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 22:18:31
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{55, 45, 36, 28, 21, 15, 10, 6, 3, 1}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] <= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}
