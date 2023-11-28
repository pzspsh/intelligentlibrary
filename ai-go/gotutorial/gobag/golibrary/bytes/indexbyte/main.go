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
	fmt.Println(bytes.IndexByte([]byte{1, 2, 3}, 3)) // 2
	fmt.Println(bytes.IndexByte([]byte{1, 2, 3}, 0)) // -1
}
