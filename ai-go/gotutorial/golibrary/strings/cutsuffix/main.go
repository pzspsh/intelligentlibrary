package main

import (
	"fmt"
	"strings"
)

func main() {
	show := func(s, sep string) {
		before, found := strings.CutSuffix(s, sep)
		fmt.Printf("CutSuffix(%q, %q) = %q, %v\n", s, sep, before, found)
	}
	show("Gopher", "Go")
	show("Gopher", "er")
}
