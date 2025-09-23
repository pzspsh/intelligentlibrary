package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("OnesCount16(%016b) = %d\n", 14, bits.OnesCount16(14))
}
