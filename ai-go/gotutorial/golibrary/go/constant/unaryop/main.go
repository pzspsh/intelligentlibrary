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
	vs := []constant.Value{
		constant.MakeBool(true),
		constant.MakeFloat64(2.7),
		constant.MakeUint64(42),
	}

	for i, v := range vs {
		switch v.Kind() {
		case constant.Bool:
			vs[i] = constant.UnaryOp(token.NOT, v, 0)

		case constant.Float:
			vs[i] = constant.UnaryOp(token.SUB, v, 0)

		case constant.Int:
			// Use 16-bit precision.
			// This would be equivalent to ^uint16(v).
			vs[i] = constant.UnaryOp(token.XOR, v, 16)
		}
	}

	for _, v := range vs {
		fmt.Println(v)
	}

}
