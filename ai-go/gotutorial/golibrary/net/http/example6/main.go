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
		fmt.Println(req.PostFormValue("a")) //获取post的a值
		req.ParseForm()
		fmt.Println(req.PostForm) //只包含请求体数据
	})
	http.ListenAndServe("127.0.0.1:9999", nil)
}
