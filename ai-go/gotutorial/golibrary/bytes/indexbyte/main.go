/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:45:29
*/
package main

import (
	"bytes"
	"fmt"
)

/*
IndexByte返回b中c的第一个实例的索引，如果b中不存在c，则返回-1。
*/
func main() {
	fmt.Println(bytes.IndexByte([]byte("chicken"), byte('k'))) // 4
	fmt.Println(bytes.IndexByte([]byte("chicken"), byte('g'))) // -1
	fmt.Println(bytes.IndexByte([]byte{1, 2, 3}, 3))           // 2
	fmt.Println(bytes.IndexByte([]byte{1, 2, 3}, 0))           // -1
}
