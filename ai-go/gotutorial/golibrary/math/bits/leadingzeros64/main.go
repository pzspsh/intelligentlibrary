package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("LeadingZeros64(%064b) = %d\n", 1, bits.LeadingZeros64(1))
}
