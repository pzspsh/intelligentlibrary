/*
@File   : main.go
@Author : pan
@Time   : 2024-08-19 17:10:14
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// 网络请求响应体互转
func main() {
	// 发起HTTP GET请求
	resp, err := http.Get("http://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 将字节切片转换回io.Reader
	reader := bytes.NewReader(body)

	// 现在你可以使用reader，‌比如再次读取数据或者传递给其他需要io.Reader的函数
	// 这里我们只是简单地打印出读取到的数据
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", data)
}
