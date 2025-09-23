package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("TrailingZeros32(%032b) = %d\n", 14, bits.TrailingZeros32(14))
}
