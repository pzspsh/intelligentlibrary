/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:56:31
*/
package main

import (
	"bytes"
	"fmt"
)

/*
Join将s中的元素连接起来以创建一个新的字节片。分隔符sep放置在结果片中的元素之间。
*/
func main() {
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("%s", bytes.Join(s, []byte(", ")))
	fmt.Println(bytes.Join([][]byte{{1, 1}, {2, 2}, {3, 3}}, []byte{9})) // [1 1 9 2 2 9 3 3]
}
