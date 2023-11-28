/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:53:22
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Replace([]byte{1, 2, 1, 2, 3, 1, 2, 1, 2}, []byte{1, 2}, []byte{0, 0}, 3)) // [0 0 0 0 3 0 0 1 2]
}
