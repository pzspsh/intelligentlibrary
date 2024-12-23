/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 14:10:26
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	f := new(big.Float)
	f.SetString("3.14159")
	fmt.Println(f)
}
