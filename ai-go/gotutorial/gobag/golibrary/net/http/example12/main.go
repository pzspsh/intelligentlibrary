/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:44:19
*/
package main

import (
	"net/http"
)

/* 文件服务器（类似Nginx） */
func main() {
	http.Handle("/static/", http.FileServer(http.Dir("./")))
	//第一种
	http.Handle("/static2/", http.StripPrefix("/static2/", http.FileServer(http.Dir("./www"))))
	//第二种
	http.Handle("/www/", http.FileServer(http.Dir(".")))
	http.ListenAndServe("127.0.0.1:9999", nil)
}
