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
	// First number is 33<<64 + 12
	n1 := []uint64{33, 12}
	// Second number is 21<<64 + 23
	n2 := []uint64{21, 23}
	// Add them together without producing carry.
	d1, carry := bits.Add64(n1[1], n2[1], 0)
	d0, _ := bits.Add64(n1[0], n2[0], carry)
	nsum := []uint64{d0, d1}
	fmt.Printf("%v + %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)

	// First number is 1<<64 + 9223372036854775808
	n1 = []uint64{1, 0x8000000000000000}
	// Second number is 1<<64 + 9223372036854775808
	n2 = []uint64{1, 0x8000000000000000}
	// Add them together producing carry.
	d1, carry = bits.Add64(n1[1], n2[1], 0)
	d0, _ = bits.Add64(n1[0], n2[0], carry)
	nsum = []uint64{d0, d1}
	fmt.Printf("%v + %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)
}
