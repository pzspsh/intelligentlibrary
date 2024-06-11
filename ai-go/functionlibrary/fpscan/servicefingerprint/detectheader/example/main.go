/*
@File   : main.go
@Author : pan
@Time   : 2024-06-11 11:19:34
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	target := fmt.Sprintf("http://%s:%v", "target_ip/domain_address", "target_port") // 替换为目标IP和端口

	resp, err := http.Get(target)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", target, err)
		return
	}
	defer resp.Body.Close()

	// 读取并打印User-Agent字段
	userAgent := resp.Header.Get("User-Agent")
	fmt.Printf("User-Agent: %s\n", userAgent)

	// 你可以进一步分析User-Agent字段，尝试提取操作系统信息
	// 但请注意，User-Agent可以被伪造，因此它可能不是完全可靠的指示器
}
