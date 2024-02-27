/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:48:37
*/
package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	handler := &HelloHandler{}
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}
