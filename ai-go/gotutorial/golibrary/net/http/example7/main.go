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

//解析参数
//http协议
//post获取数据

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		fmt.Println(req.Form)
		fmt.Println(req.Form.Get("a"))  //获取a的值
		fmt.Println(req.Form["a"])      //获取a的值
		fmt.Println(req.FormValue("a")) //获取a的值
		fmt.Println(req.PostForm)       //只包含请求体数据
	})
	http.ListenAndServe("127.0.0.1:9999", nil)
}
