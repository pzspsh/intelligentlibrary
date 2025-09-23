package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("LeadingZeros32(%032b) = %d\n", 1, bits.LeadingZeros32(1))
}
