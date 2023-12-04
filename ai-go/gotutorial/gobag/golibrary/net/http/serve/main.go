/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:37:38
*/
package main

import (
	"fmt"
	"net"
	"net/http"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Gophers!")
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	handler := &myHandler{}

	fmt.Println("Starting server on :8080")
	if err := http.Serve(listener, handler); err != nil {
		panic(err)
	}
}
