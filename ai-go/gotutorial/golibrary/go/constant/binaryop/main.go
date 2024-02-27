/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:46:22
*/
package main

import (
	"fmt"
	"go/constant"
	"go/token"
)

func main() {
	// 11 / 0.5
	a := constant.MakeUint64(11)
	b := constant.MakeFloat64(0.5)
	c := constant.BinaryOp(a, token.QUO, b)
	fmt.Println(c)

}
