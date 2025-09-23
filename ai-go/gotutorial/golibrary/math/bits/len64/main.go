package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("Len64(%064b) = %d\n", 8, bits.Len64(8))
}
