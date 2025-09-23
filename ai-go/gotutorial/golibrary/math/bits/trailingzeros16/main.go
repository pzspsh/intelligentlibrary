package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("TrailingZeros16(%016b) = %d\n", 14, bits.TrailingZeros16(14))
}
