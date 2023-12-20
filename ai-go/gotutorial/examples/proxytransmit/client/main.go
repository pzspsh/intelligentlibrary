/*
@File   : main.go
@Author : pan
@Time   : 2023-12-20 17:56:32
*/
package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("topgoer.com是个不错的网站我是9091里面的")
}
func main() {
	http.HandleFunc("/topgoer", sayHello)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}

// 访问http://localhost:9090/5lmh实际执行的是http://localhost:9091/5lmh的方法
