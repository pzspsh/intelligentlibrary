package main

import (
	"fmt"
	"net/http"
)

func main() {
	m, n, ok := http.ParseHTTPVersion("HTTP/1.0")
	fmt.Println(m, n, ok) //1 0 true
}
