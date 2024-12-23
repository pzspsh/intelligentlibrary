/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 13:54:32
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.6f\n", math.Expm1(0.01))
	fmt.Printf("%.6f\n", math.Expm1(-1))
}
