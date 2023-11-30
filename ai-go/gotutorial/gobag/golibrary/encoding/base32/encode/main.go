/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:46:13
*/
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	data := []byte("Hello, world!")
	dst := make([]byte, base32.StdEncoding.EncodedLen(len(data)))
	base32.StdEncoding.Encode(dst, data)
	fmt.Println(string(dst))
}
