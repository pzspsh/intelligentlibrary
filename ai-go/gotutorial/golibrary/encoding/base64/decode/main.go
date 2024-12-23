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
	str := "SGVsbG8sIHdvcmxkIQ=="
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(str)))
	n, err := base64.StdEncoding.Decode(dst, []byte(str))
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	dst = dst[:n]
	fmt.Printf("%q\n", dst)
}
