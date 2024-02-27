/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 12:34:59
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// 公共打印方法
func printBody(r *http.Response) {
	defer func() { _ = r.Body.Close() }()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

// 设置请求参数
func requestByParams() {
	//创建一个请求
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	//添加请求参数
	params := make(url.Values)
	params.Add("name", "kimi")
	params.Add("age", "18")
	//将参数绑定到request上,params.Encode()就等于name=kimi&age=18
	request.URL.RawQuery = params.Encode()

	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	printBody(r)

}

func main() {
	//如何设置请求的查询参数，http://httpbin.org/get?name=kimi&age=18
	requestByParams()
}
