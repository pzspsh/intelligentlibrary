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
Index返回s中sep的第一个实例的索引，如果s中不存在sep，则返回-1。
*/
func main() {
	fmt.Println(bytes.Index([]byte("chicken"), []byte("ken")))    // 4
	fmt.Println(bytes.Index([]byte("chicken"), []byte("dmr")))    // -1
	fmt.Println(bytes.Index([]byte{1, 2, 3, 4, 5}, []byte{4, 5})) // 3
	fmt.Println(bytes.Index([]byte{1, 2, 3, 4, 5}, []byte{0, 1})) // -1
}
