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
	fmt.Println(bytes.Index([]byte{1, 2, 3, 4, 5}, []byte{4, 5})) // 3
	fmt.Println(bytes.Index([]byte{1, 2, 3, 4, 5}, []byte{0, 1})) // -1
}
