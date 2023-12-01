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
	// First number is 0<<32 + 12
	n1 := []uint32{0, 12}
	// Second number is 0<<32 + 12
	n2 := []uint32{0, 12}
	// Multiply them together without producing overflow.
	hi, lo := bits.Mul32(n1[1], n2[1])
	nsum := []uint32{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)

	// First number is 0<<32 + 2147483648
	n1 = []uint32{0, 0x80000000}
	// Second number is 0<<32 + 2
	n2 = []uint32{0, 2}
	// Multiply them together producing overflow.
	hi, lo = bits.Mul32(n1[1], n2[1])
	nsum = []uint32{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)
}
