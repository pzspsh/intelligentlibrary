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
	i := new(big.Int)
	i.SetString("644", 8) // octal
	fmt.Println(i)
}
