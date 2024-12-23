package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(multiply1("123", "456"))
}

func multiply1(num1 string, num2 string) string {
	a, b := new(big.Int), new(big.Int)
	a.SetString(num1, 10)
	b.SetString(num2, 10)
	a.Mul(a, b)
	return a.String()
}
