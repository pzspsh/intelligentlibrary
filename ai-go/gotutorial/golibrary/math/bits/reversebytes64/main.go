package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%064b\n", 15)
	fmt.Printf("%064b\n", bits.ReverseBytes64(15))
}
