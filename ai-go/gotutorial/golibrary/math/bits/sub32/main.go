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
	// First number is 33<<32 + 23
	n1 := []uint32{33, 23}
	// Second number is 21<<32 + 12
	n2 := []uint32{21, 12}
	// Sub them together without producing carry.
	d1, carry := bits.Sub32(n1[1], n2[1], 0)
	d0, _ := bits.Sub32(n1[0], n2[0], carry)
	nsum := []uint32{d0, d1}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)

	// First number is 3<<32 + 2147483647
	n1 = []uint32{3, 0x7fffffff}
	// Second number is 1<<32 + 2147483648
	n2 = []uint32{1, 0x80000000}
	// Sub them together producing carry.
	d1, carry = bits.Sub32(n1[1], n2[1], 0)
	d0, _ = bits.Sub32(n1[0], n2[0], carry)
	nsum = []uint32{d0, d1}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)
}
