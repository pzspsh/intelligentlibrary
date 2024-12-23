/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 14:55:08
*/
package main

import (
	"bytes"
	"fmt"
)

/*
SplitAfter在每个sep实例之后将s切片为所有子切片，并返回这些子切片的一个切片。
如果sep为空，SplitAfter在每个UTF-8序列之后进行分割。它相当于计数为-1的SplitAfterN。
*/
func main() {
	fmt.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c"), []byte(","))) // ["a," "b," "c"]
}
