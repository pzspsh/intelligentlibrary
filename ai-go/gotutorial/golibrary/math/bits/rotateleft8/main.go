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
	fmt.Printf("%08b\n", 15)
	fmt.Printf("%08b\n", bits.RotateLeft8(15, 2))
	fmt.Printf("%08b\n", bits.RotateLeft8(15, -2))
}
