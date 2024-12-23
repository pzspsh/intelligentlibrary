/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:20:12
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

/*
NewRequest中body参数对req.Body、req.ContentLength值的影响
*/
func main() {
	readByte := func(r io.Reader) io.Reader {
		var b [1]byte
		r.Read(b[:])
		return r
	}
	tests := []struct {
		r    io.Reader
		want int64
	}{
		{bytes.NewReader([]byte("123")), 3},
		{bytes.NewBuffer([]byte("1234")), 4},
		{strings.NewReader("12345"), 5},
		{strings.NewReader(""), 0},

		// Not detected. During Go 1.8 we tried to make these set to -1, but
		// due to Issue 18117, we keep these returning 0, even though they're
		// unknown.
		{struct{ io.Reader }{strings.NewReader("xyz")}, 0},
		{io.NewSectionReader(strings.NewReader("x"), 0, 6), 0},
		{readByte(io.NewSectionReader(strings.NewReader("xy"), 0, 6)), 0},
	}
	for i, tt := range tests {
		req, err := http.NewRequest("POST", "http://localhost/", tt.r)
		fmt.Println(req.Body)
		if err != nil {
			fmt.Println(err)
		}
		if req.ContentLength != tt.want { //没有返回，说明req.ContentLength == tt.want
			fmt.Printf("test[%d]: ContentLength(%T) = %d; want %d", i, tt.r, req.ContentLength, tt.want)
		}
	}
}
