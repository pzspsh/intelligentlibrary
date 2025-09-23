package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	fmt.Printf("%.1f", cmplx.Exp(1i*math.Pi)+1)
}
