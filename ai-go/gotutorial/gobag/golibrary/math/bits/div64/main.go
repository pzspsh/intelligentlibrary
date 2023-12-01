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
	// First number is 0<<64 + 6
	n1 := []uint64{0, 6}
	// Second number is 0<<64 + 3
	n2 := []uint64{0, 3}
	// Divide them together.
	quo, rem := bits.Div64(n1[0], n1[1], n2[1])
	nsum := []uint64{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)

	// First number is 2<<64 + 9223372036854775808
	n1 = []uint64{2, 0x8000000000000000}
	// Second number is 0<<64 + 9223372036854775808
	n2 = []uint64{0, 0x8000000000000000}
	// Divide them together.
	quo, rem = bits.Div64(n1[0], n1[1], n2[1])
	nsum = []uint64{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)
}
