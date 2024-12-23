/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:39:57
*/
package main

import (
	"fmt"
	"net/http"
)

type MyMux struct{}

func (mux *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhello(w, r)
	}
	http.NotFound(w, r)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
