package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"alice", "Bob", "VERA"}
	slices.Reverse(names)
	fmt.Println(names)
}
