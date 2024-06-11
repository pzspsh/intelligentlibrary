/*
@File   : main.go
@Author : pan
@Time   : 2024-06-11 11:06:21
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	target := fmt.Sprintf("http://%v:%v", "127.0.0.1", "80") // 替换为目标IP和端口

	resp, err := http.Get(target)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", target, err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// 输出响应头和响应体内容
	fmt.Printf("Status: %s\n", resp.Status)
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
	fmt.Printf("Body: %s\n", string(body))
}
