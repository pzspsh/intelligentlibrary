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

func main() {
	fmt.Println(bytes.LastIndex([]byte("hi go"), []byte("go"))) // 3
	fmt.Println(bytes.LastIndex([]byte{1, 2, 3}, []byte{2, 3})) // 1
}
