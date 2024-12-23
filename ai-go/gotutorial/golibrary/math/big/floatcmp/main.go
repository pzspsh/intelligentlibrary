/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 14:10:26
*/
package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	inf := math.Inf(1)
	zero := 0.0

	operands := []float64{-inf, -1.2, -zero, 0, +1.2, +inf}

	fmt.Println("   x     y  cmp")
	fmt.Println("---------------")
	for _, x64 := range operands {
		x := big.NewFloat(x64)
		for _, y64 := range operands {
			y := big.NewFloat(y64)
			fmt.Printf("%4g  %4g  %3d\n", x, y, x.Cmp(y))
		}
		fmt.Println()
	}

}
