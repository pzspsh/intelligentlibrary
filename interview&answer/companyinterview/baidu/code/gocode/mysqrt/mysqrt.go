/*
@File   : mysqrt.go
@Author : pan
@Time   : 2023-05-23 10:12:26
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(10))
}

// x的平方根
func mySqrt(x int) int {
	result := x
	for result*result > x {
		result = (result + x/result) / 2
	}
	return result
}
