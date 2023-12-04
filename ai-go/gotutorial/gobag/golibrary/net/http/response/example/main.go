/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:36:02
*/
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	r := &http.Response{
		Status:     "123 some status",
		StatusCode: 123,
		ProtoMajor: 1,
		ProtoMinor: 3,
	}
	var buf bytes.Buffer
	r.Write(&buf)
	fmt.Println(buf.String())
	if strings.Contains(buf.String(), "123 123") {
		fmt.Printf("stutter in status: %s", buf.String())
	}
}
