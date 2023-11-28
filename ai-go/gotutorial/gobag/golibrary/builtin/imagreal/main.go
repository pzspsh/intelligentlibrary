/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:51:44
*/
package main

import (
	"fmt"
)

func main() {
	a := complex(1, 2)
	b := imag(a)
	c := real(a)
	fmt.Println(a, b, c)
}
