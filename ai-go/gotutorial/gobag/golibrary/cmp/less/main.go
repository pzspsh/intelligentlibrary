/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:20:26
*/
package main

import (
	"cmp"
	"fmt"
	"math"
)

/*
Less报告x是否小于y。对于浮点类型，NaN被认为小于任何非NaN，并且-0.0不小于(等于)0.0。
*/
func main() {
	fmt.Println(cmp.Less(1, 2))                   //true
	fmt.Println(cmp.Less(1, 1))                   // false
	fmt.Println(cmp.Less(2, 1))                   // false
	fmt.Println(cmp.Less(1, math.NaN()))          // false
	fmt.Println(cmp.Less(-1, math.NaN()))         // false
	fmt.Println(cmp.Less(math.NaN(), math.NaN())) // false
	fmt.Println(cmp.Less(math.NaN(), -1.0))       // true
	// fmt.Println(cmp.Less(-0.0, 0.0))              // false
}
