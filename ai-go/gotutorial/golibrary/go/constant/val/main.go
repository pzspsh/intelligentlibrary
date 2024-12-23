/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:46:22
*/
package main

import (
	"fmt"
	"go/constant"
	"math"
)

func main() {
	maxint := constant.MakeInt64(math.MaxInt64)
	fmt.Printf("%v\n", constant.Val(maxint))

	e := constant.MakeFloat64(math.E)
	fmt.Printf("%v\n", constant.Val(e))

	b := constant.MakeBool(true)
	fmt.Printf("%v\n", constant.Val(b))

	b = constant.Make(false)
	fmt.Printf("%v\n", constant.Val(b))

}
