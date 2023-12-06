/*
@File   : main.go
@Author : pan
@Time   : 2023-12-06 10:26:43
*/
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!")
}

/*
实现一个最简单的HTTPS Web Server
*/
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
}
