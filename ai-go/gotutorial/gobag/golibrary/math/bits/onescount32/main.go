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
	fmt.Printf("OnesCount32(%032b) = %d\n", 14, bits.OnesCount32(14))
}
