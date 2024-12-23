/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 14:10:26
*/
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	r, theta := cmplx.Polar(2i)
	fmt.Printf("r: %.1f, θ: %.1f*π", r, theta/math.Pi)
}
