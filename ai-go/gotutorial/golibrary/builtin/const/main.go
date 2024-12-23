/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:48:36
*/
package main

import (
	"fmt"
)

const (
	a = iota //iota默认初始值为０
	b = 100
	c        //c默认跟上一个赋值相同
	d = iota //iota默认每行加１，故此时其值为３
)

func main() {
	fmt.Println(a, b, c, d)
}
