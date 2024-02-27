/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:44:19
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home/", func(resp http.ResponseWriter, req *http.Request) {
		//fmt.Fprint(resp, "首页")//当前页面
		//进行重镜像
		http.Redirect(resp, req, "/Login/", http.StatusFound) //将网址重定向到login中
	})
	http.HandleFunc("/Login/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprint(resp, "登录")
	})
	http.ListenAndServe("127.0.0.1:9999", nil)
}
