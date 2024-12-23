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
	fmt.Printf("OnesCount8(%08b) = %d\n", 14, bits.OnesCount8(14))
}
