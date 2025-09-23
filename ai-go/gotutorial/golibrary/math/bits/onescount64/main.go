package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("OnesCount64(%064b) = %d\n", 14, bits.OnesCount64(14))
}
