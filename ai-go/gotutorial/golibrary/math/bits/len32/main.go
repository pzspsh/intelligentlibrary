package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("Len32(%032b) = %d\n", 8, bits.Len32(8))
}
