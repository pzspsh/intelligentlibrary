package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%032b\n", 15)
	fmt.Printf("%032b\n", bits.RotateLeft32(15, 2))
	fmt.Printf("%032b\n", bits.RotateLeft32(15, -2))
}
