package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("TrailingZeros8(%08b) = %d\n", 14, bits.TrailingZeros8(14))
}
