package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%032b\n", 19)
	fmt.Printf("%032b\n", bits.Reverse32(19))
}
