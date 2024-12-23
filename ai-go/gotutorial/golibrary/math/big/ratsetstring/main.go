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
	r := new(big.Rat)
	r.SetString("355/113")
	fmt.Println(r.FloatString(3))
}
