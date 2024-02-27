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
	// First number is 0<<64 + 12
	n1 := []uint64{0, 12}
	// Second number is 0<<64 + 12
	n2 := []uint64{0, 12}
	// Multiply them together without producing overflow.
	hi, lo := bits.Mul64(n1[1], n2[1])
	nsum := []uint64{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)

	// First number is 0<<64 + 9223372036854775808
	n1 = []uint64{0, 0x8000000000000000}
	// Second number is 0<<64 + 2
	n2 = []uint64{0, 2}
	// Multiply them together producing overflow.
	hi, lo = bits.Mul64(n1[1], n2[1])
	nsum = []uint64{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)
}
