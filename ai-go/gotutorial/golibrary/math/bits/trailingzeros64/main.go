package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("TrailingZeros64(%064b) = %d\n", 14, bits.TrailingZeros64(14))
}
