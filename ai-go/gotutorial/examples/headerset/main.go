/*
@File   : main.go
@Author : pan
@Time   : 2024-08-22 11:04:59
*/
package main

import (
	"fmt"
	"net/http"
)

func HeaderSet() {
	url := "http://example.com"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// 设置自定义headers
	req.Header.Set("Custom-Header", "HeaderValue")
	req.Header.Set("Another-Header", "AnotherValue")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

func HeaderAdd() {
	url := "http://example.com"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 添加自定义的headers
	req.Header.Add("Custom-Header", "HeaderValue")
	req.Header.Add("Another-Header", "AnotherValue")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 处理响应...
	fmt.Println("Response Status:", resp.Status)
}

func main() {
	// http请求header头的设置和添加
}
