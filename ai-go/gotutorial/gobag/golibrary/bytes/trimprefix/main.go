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
	fmt.Println(string(bytes.TrimPrefix([]byte("hi hi go"), []byte("hi")))) //  hi go
}
