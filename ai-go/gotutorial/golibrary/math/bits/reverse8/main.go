package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%08b\n", 19)
	fmt.Printf("%08b\n", bits.Reverse8(19))
}
