/*
@File   : main.go
@Author : pan
@Time   : 2024-06-11 11:48:25
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	target := "http://example.com/some_specific_path" // 替换为目标IP、端口和特定路径

	resp, err := http.Get(target)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", target, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// 分析响应体，查找特定于操作系统的错误消息或模式
	// 这通常需要对目标服务及其在不同操作系统上的行为有深入了解
	fmt.Printf("Response body: %s\n", string(body))
}
