package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("Len8(%08b) = %d\n", 8, bits.Len8(8))
}
