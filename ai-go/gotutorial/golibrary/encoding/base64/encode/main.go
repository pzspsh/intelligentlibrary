/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:53:51
*/
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, world!")
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)
	fmt.Println(string(dst))
}
