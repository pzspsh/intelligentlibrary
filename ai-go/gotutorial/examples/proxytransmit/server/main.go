/*
@File   : main.go
@Author : pan
@Time   : 2023-12-20 17:56:32
*/
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse("http://127.0.0.1:9091/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}
func main() {
	http.HandleFunc("/topgoer", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}
