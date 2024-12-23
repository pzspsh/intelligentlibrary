/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:33:01
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	req := &http.Request{
		Method: "POST",
		Header: http.Header{"Content-Type": {`multipart/form-data; boundary="foo123"`}},
		Body:   io.NopCloser(new(bytes.Buffer)),
	}
	multipart, err := req.MultipartReader() //r.Body将作为流处理
	if multipart == nil {
		fmt.Printf("expected multipart; error: %v", err)
	}
	fmt.Println(multipart) //&{0xc00007e240 <nil> 0 [13 10] [13 10 45 45 102 111 111 49 50 51] [45 45 102 111 111 49 50 51 45 45] [45 45 102 111 111 49 50 51]}

	req = &http.Request{
		Method: "POST",
		Header: http.Header{"Content-Type": {`multipart/mixed; boundary="foo123"`}},
		Body:   io.NopCloser(new(bytes.Buffer)),
	}
	multipart, err = req.MultipartReader()
	if multipart == nil {
		fmt.Printf("expected multipart; error: %v", err)
	}
	fmt.Println(multipart) //&{0xc00007e2a0 <nil> 0 [13 10] [13 10 45 45 102 111 111 49 50 51] [45 45 102 111 111 49 50 51 45 45] [45 45 102 111 111 49 50 51]}

	req.Header = http.Header{"Content-Type": {"text/plain"}}
	multipart, err = req.MultipartReader()
	if err != nil {
		fmt.Println(err)
	}
	if multipart != nil {
		fmt.Printf("unexpected multipart for text/plain")
	}
	fmt.Println(multipart) //<nil>
}
