package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mySqrt1(10))
	fmt.Println(mySqrt1(1))
}

func mySqrt1(x int) int {
	result := math.Floor(math.Sqrt(float64(x)))
	return int(result)
}
