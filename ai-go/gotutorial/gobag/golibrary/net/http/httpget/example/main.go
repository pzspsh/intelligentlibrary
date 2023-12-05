/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:02:55
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// 启动server/main.go端
func main() {
	apiUrl := "http://127.0.0.1:8080/get"
	// URL param
	data := url.Values{}
	data.Set("name", "王二小")
	data.Set("age", "18")
	//把string转换为url
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	}
	//将参数添加到请求url
	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
