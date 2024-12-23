/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:35:25
*/
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" //注意下划线
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	go func() {
		fmt.Println("pprof start...")
		fmt.Println(http.ListenAndServe(":9998", nil))
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", SayHello)

	err := http.ListenAndServe(":9999", mux)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
