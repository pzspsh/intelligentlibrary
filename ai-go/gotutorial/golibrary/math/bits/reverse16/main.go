package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%016b\n", 19)
	fmt.Printf("%016b\n", bits.Reverse16(19))
}
