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
	p := math.Round(10.5)
	fmt.Printf("%.1f\n", p)

	n := math.Round(-10.5)
	fmt.Printf("%.1f\n", n)
}
