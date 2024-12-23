/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:19:26
*/
package main

import (
	"cmp"
	"fmt"
	"math"
)

func main() {
	fmt.Println(cmp.Compare(1, 2))                   // -1
	fmt.Println(cmp.Compare(1, 1))                   // 0
	fmt.Println(cmp.Compare(2, 1))                   // 1
	fmt.Println(cmp.Compare(1, math.NaN()))          // 1
	fmt.Println(cmp.Compare(-1, math.NaN()))         // 1
	fmt.Println(cmp.Compare(math.NaN(), math.NaN())) // 0
	// fmt.Println(cmp.Compare(-0.0, 0.0))              // 0
}

/*
-1 if x is less than y,
0 if x equals y,
+1 if x is greater than y.

如果x小于y -1，
x = y时为0，
x大于y +1。
*/
