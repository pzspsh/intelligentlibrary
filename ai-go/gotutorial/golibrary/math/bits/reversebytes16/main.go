package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Printf("%016b\n", 15)
	fmt.Printf("%016b\n", bits.ReverseBytes16(15))
}
