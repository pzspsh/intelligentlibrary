package main

import (
	"fmt"
)

const (
	a    = iota
	b, c = iota, iota
	d    = iota
)

func main() {
	fmt.Println(a, b, c, d)
}
