/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:19:15
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//打印请求主机地址
	fmt.Println(r.Host)
	//打印请求头信息
	fmt.Printf("header content:[%v]\n", r.Header)

	//获取 post 请求中 form 里边的数据
	fmt.Printf("form content:[%s, %s]\n", r.PostFormValue("username"), r.PostFormValue("passwd"))

	//读取请求体信息
	bodyContent, err := io.ReadAll(r.Body)
	if err != nil && err != io.EOF {
		fmt.Printf("read body content failed, err:[%s]\n", err.Error())
		return
	}
	fmt.Printf("body content:[%s]\n", string(bodyContent))

	//返回响应内容
	fmt.Fprintf(w, "hello world ~")
}

func main() {
	http.HandleFunc("/index", IndexHandler)
	http.ListenAndServe("10.10.19.200:8000", nil)
}
