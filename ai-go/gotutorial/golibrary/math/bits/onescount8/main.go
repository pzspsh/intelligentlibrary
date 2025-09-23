package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("OnesCount8(%08b) = %d\n", 14, bits.OnesCount8(14))
}
