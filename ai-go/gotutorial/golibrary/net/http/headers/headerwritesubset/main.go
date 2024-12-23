/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:15:24
*/
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

var headerWriteTests = []struct {
	h        http.Header
	exclude  map[string]bool
	expected string
}{
	{http.Header{}, nil, ""},
	{
		http.Header{
			"Content-Type":   {"text/html; charset=UTF-8"},
			"Content-Length": {"0"},
		},
		nil,
		"Content-Length: 0\r\nContent-Type: text/html; charset=UTF-8\r\n",
	},
	{
		http.Header{
			"Expires":          {"-1"},
			"Content-Length":   {"0"},
			"Content-Encoding": {"gzip"},
		},
		map[string]bool{"Content-Length": true}, //"Content-Length"字段将不会写入io.Writer
		"Content-Encoding: gzip\r\nExpires: -1\r\n",
	},
}

func main() {
	var buf bytes.Buffer //得到io.Writer
	for i, test := range headerWriteTests {
		test.h.WriteSubset(&buf, test.exclude)
		fmt.Println(i)
		buf.WriteTo(os.Stdout)
		fmt.Println()
		if buf.String() != test.expected {
			fmt.Printf("#%d:\n got: %q\nwant: %q", i, buf.String(), test.expected)
		}
		buf.Reset()
	}
}
