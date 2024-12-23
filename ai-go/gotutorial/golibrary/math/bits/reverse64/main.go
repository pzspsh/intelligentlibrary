/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 14:10:26
*/
package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%064b\n", 19)
	fmt.Printf("%064b\n", bits.Reverse64(19))
}
