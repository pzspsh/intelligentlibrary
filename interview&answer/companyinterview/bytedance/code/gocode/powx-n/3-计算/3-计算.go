package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myPow2(2.00000, 10))
}

func myPow2(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
