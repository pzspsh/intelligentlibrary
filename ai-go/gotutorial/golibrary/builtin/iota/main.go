/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:49:39
*/
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
