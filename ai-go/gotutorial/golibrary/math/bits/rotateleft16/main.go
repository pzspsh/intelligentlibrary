package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%016b\n", 15)
	fmt.Printf("%016b\n", bits.RotateLeft16(15, 2))
	fmt.Printf("%016b\n", bits.RotateLeft16(15, -2))
}
