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
	sin, cos := math.Sincos(0)
	fmt.Printf("%.2f, %.2f", sin, cos)
}
