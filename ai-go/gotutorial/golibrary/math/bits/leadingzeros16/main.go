package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("LeadingZeros16(%016b) = %d\n", 1, bits.LeadingZeros16(1))
}
