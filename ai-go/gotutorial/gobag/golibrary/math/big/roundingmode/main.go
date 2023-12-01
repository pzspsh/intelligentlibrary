/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 14:10:26
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	operands := []float64{2.6, 2.5, 2.1, -2.1, -2.5, -2.6}

	fmt.Print("   x")
	for mode := big.ToNearestEven; mode <= big.ToPositiveInf; mode++ {
		fmt.Printf("  %s", mode)
	}
	fmt.Println()

	for _, f64 := range operands {
		fmt.Printf("%4g", f64)
		for mode := big.ToNearestEven; mode <= big.ToPositiveInf; mode++ {
			// sample operands above require 2 bits to represent mantissa
			// set binary precision to 2 to round them to integer values
			f := new(big.Float).SetPrec(2).SetMode(mode).SetFloat64(f64)
			fmt.Printf("  %*g", len(mode.String()), f)
		}
		fmt.Println()
	}

}
