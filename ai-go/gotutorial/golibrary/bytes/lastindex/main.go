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
LastIndex返回s中sep的最后一个实例的索引，如果s中不存在sep，则返回-1。
*/
func main() {
	fmt.Println(bytes.Index([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("rodent")))
	fmt.Println(bytes.LastIndex([]byte("hi go"), []byte("go"))) // 3
	fmt.Println(bytes.LastIndex([]byte{1, 2, 3}, []byte{2, 3})) // 1
}
