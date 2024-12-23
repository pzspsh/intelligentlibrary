/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:44:27
*/
package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

func main() {
	var in bytes.Buffer
	b := []byte(`xiorui.cc`)
	w := zlib.NewWriter(&in)
	w.Write(b)
	w.Close()

	var out bytes.Buffer
	r, _ := zlib.NewReader(&in)
	io.Copy(&out, r)
	fmt.Println(out.String())

}
