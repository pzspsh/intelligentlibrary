package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("OnesCount32(%032b) = %d\n", 14, bits.OnesCount32(14))
}
