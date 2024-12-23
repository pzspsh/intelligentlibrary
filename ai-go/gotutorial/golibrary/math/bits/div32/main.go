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
	// First number is 0<<32 + 6
	n1 := []uint32{0, 6}
	// Second number is 0<<32 + 3
	n2 := []uint32{0, 3}
	// Divide them together.
	quo, rem := bits.Div32(n1[0], n1[1], n2[1])
	nsum := []uint32{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)

	// First number is 2<<32 + 2147483648
	n1 = []uint32{2, 0x80000000}
	// Second number is 0<<32 + 2147483648
	n2 = []uint32{0, 0x80000000}
	// Divide them together.
	quo, rem = bits.Div32(n1[0], n1[1], n2[1])
	nsum = []uint32{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)
}
