package main

import (
	"fmt"
	"math"
)

func main() {
	sin, cos := math.Sincos(0)
	fmt.Printf("%.2f, %.2f", sin, cos)
}
