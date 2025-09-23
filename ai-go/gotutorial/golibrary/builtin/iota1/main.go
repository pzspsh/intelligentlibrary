package main

import (
	"fmt"
)

const (
	a    = iota
	b, c = iota, iota
	d    = iota
)
const e = iota //再次遇到const关键字，iota值变为０

func main() {
	fmt.Println(a, b, c, d, e)
}
