/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:54:58
*/
package main

import (
	"fmt"
)

func main() {

	a := make([]int, 1)
	b := []int{1, 2}
	//由于a的大小为１，所以b只给a复制了１个元素１，并且返回ｃ的长度也是１
	c := copy(a, b)
	fmt.Println(a, b, c)
}
