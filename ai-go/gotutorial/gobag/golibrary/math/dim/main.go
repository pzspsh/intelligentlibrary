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
	fmt.Printf("%.2f\n", math.Dim(4, -2))
	fmt.Printf("%.2f\n", math.Dim(-4, 2))
}
