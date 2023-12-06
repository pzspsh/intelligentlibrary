/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 22:32:49
*/
package main

import (
	"fmt"
	"strconv"
)

/*
ParseInt解释给定基数(0,2到36)和位大小(0到64)的字符串s，并返回相应的值i。
*/
func main() {
	v32 := "-354634382"
	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
		fmt.Printf("10v32: %T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
		fmt.Printf("16v32: %T, %v\n", s, s)
	}

	v64 := "-3546343826724305832"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("10v64: %T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
		fmt.Printf("16v64: %T, %v\n", s, s)
	}
}
