/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:12:56
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	m, n, ok := http.ParseHTTPVersion("HTTP/1.0")
	fmt.Println(m, n, ok) //1 0 true
}
