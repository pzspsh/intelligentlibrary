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
	fmt.Println(bytes.Repeat([]byte{1, 2}, 3)) // [1 2 1 2 1 2]
}
