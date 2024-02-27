/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 22:03:25
*/
package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	names := []string{"bob", "Bob", "alice", "Vera", "VERA"}
	names = slices.CompactFunc(names, func(a, b string) bool {
		return strings.ToLower(a) == strings.ToLower(b)
	})
	fmt.Println(names)
}
