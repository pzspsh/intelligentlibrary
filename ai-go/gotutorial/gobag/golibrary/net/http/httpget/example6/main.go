/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 13:00:11
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetHttp(url string) (body []byte, err error) {
	// 创建 client 和 resp 对象
	var client http.Client
	var resp *http.Response

	// 这里博主设置了10秒钟的超时
	client = http.Client{Timeout: 10 * time.Second}

	// 这里使用了 Get 方法，并判断异常
	resp, err = client.Get(url)
	if err != nil {
		return nil, err
	}
	// 释放对象
	defer resp.Body.Close()

	// 把获取到的页面作为返回值返回
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// 释放对象
	defer client.CloseIdleConnections()

	return body, nil
}

func main() {
	resp, err := GetHttp("http://10.10.102.91:10251/metrics")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(resp))
}
