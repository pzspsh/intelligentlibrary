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
