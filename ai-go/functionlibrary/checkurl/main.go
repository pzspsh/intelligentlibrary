/*
@File   : main.go
@Author : pan
@Time   : 2024-07-26 14:10:31
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

// 判断怎么知道给一个域名能被http还是https访问或者两者都可以访问
// 检查URL是否可访问
func checkURL(url string) bool {
	client := &http.Client{
		Timeout: 5 * time.Second, // 设置超时时间
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("创建请求时出错: %v\n", err)
		return false
	}

	// 可选：设置请求头，如User-Agent等
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("请求 %s 失败: %v\n", url, err)
		return false
	}
	defer resp.Body.Close()

	// 检查HTTP状态码
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Printf("%s 可访问\n", url)
		return true
	}

	fmt.Printf("%s 不可访问，状态码: %d\n", url, resp.StatusCode)
	return false
}

func main() {
	domain := "example.com"

	// 检查HTTP
	httpURL := fmt.Sprintf("http://%s", domain)
	if checkURL(httpURL) {
		// HTTP可访问
		fmt.Println("http")
	}

	// 检查HTTPS
	httpsURL := fmt.Sprintf("https://%s", domain)
	if checkURL(httpsURL) {
		// HTTPS可访问
		fmt.Println("https")
	}
}
