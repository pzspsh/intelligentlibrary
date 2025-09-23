package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%064b\n", 19)
	fmt.Printf("%064b\n", bits.Reverse64(19))
}
