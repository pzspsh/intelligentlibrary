/*
@File   : stringtoreadcloser.go
@Author : pan
@Time   : 2023-06-21 15:26:46
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func BytesNewReader(text string) (io.Reader, string) {
	reader := io.NopCloser(bytes.NewReader([]byte("hello world")))
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	reader.Close()
	rawdata := buf.String()
	return reader, rawdata
}

func StringsNewReader(text string) (io.Reader, string) {
	reader := io.NopCloser(strings.NewReader("hello world"))
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	reader.Close()
	rawdata := buf.String()
	return reader, rawdata
}

func main() {
	reader, rawdata := StringsNewReader("hello world")
	fmt.Println("io.Reader数据", reader)
	fmt.Println("原数据：", rawdata)
}
