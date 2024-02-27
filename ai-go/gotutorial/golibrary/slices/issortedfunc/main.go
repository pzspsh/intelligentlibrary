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
	"strings"
)

func main() {
	names := []string{"alice", "Bob", "VERA"}
	isSortedInsensitive := slices.IsSortedFunc(names, func(a, b string) int {
		return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	fmt.Println(isSortedInsensitive)
	fmt.Println(slices.IsSorted(names))
}
