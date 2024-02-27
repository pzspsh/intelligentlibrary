/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 14:52:22
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 产生一个随机整数
	var num int64 = rand.Int63()
	fmt.Println(num)
}
