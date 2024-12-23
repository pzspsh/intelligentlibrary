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
HasSuffix测试字节片s是否以后缀结尾。
*/
func main() {
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("go")))    // true
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("O")))     // false
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("Ami")))   // false
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("")))      // true
	fmt.Println(bytes.HasSuffix([]byte{1, 2, 3, 3}, []byte{3, 3})) // true
	fmt.Println(bytes.HasSuffix([]byte{1, 2, 3, 3}, []byte{3, 4})) // false
}
