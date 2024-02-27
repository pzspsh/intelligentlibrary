/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:59:39
*/
package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
)

func main() {
	var b bytes.Buffer

	w := zlib.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()
	fmt.Println(b.Bytes())
}
