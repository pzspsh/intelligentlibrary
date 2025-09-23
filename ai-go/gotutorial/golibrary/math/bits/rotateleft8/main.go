package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%08b\n", 15)
	fmt.Printf("%08b\n", bits.RotateLeft8(15, 2))
	fmt.Printf("%08b\n", bits.RotateLeft8(15, -2))
}
