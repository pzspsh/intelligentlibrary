/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:59:04
*/
package main

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"
)

func main() {
	buff := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207,
		47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
	b := bytes.NewReader(buff)

	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, r)

	r.Close()
}
