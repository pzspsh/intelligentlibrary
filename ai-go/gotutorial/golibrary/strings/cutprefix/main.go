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
