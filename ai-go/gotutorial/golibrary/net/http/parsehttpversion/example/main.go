/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:41:49
*/
package main

import (
	"fmt"

	"net/http"
)

func main() {
	httpVersion := "HTTP/1.1"
	major, minor, ok := http.ParseHTTPVersion(httpVersion)
	if ok {
		fmt.Printf("Major Version: %d\n", major)
		fmt.Printf("Minor Version: %d\n", minor)

	} else {
		fmt.Println("Invalid HTTP version string")
	}
}
