/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 23:08:40
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	show := func(s, sep string) {
		after, found := strings.CutPrefix(s, sep)
		fmt.Printf("CutPrefix(%q, %q) = %q, %v\n", s, sep, after, found)
	}
	show("Gopher", "Go")
	show("Gopher", "ph")
}
