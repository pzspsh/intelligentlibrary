/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:15:32
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.NewReader([]byte("Hi!")).Len())    // 3
	fmt.Println(bytes.NewReader([]byte("こんにちは!")).Len()) // 16
}
