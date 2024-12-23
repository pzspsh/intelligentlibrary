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
	const (
		a = 3
		b = 4
	)
	c := math.Sqrt(a*a + b*b)
	fmt.Printf("%.1f", c)
}
