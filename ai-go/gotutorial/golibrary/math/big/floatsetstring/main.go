package main

import (
	"fmt"
	"math/big"
)

func main() {
	f := new(big.Float)
	f.SetString("3.14159")
	fmt.Println(f)
}
