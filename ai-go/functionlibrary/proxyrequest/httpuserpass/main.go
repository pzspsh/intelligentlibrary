/*
@File   : main.go
@Author : pan
@Time   : 2024-08-01 11:18:43
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// 代理服务器的URL，包括用户名和密码（如果代理支持这种格式）
	proxyURL := "http://username:password@yourproxy.com:8080"

	// 创建一个http.Client，使用http.ProxyURL来设置代理
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(mustParseURL(proxyURL)),
		},
	}

	// 发起请求
	resp, err := client.Get("http://example.com")
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Println(string(body))
}

// mustParseURL 解析URL，并在失败时panic
func mustParseURL(rawurl string) *url.URL {
	u, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}
	return u
}

// 注意：上面的代码中，mustParseURL函数用于确保URL被正确解析。
// 然而，Go的标准库可能不支持直接从URL中解析出代理的认证信息。
// 如果代理服务器不支持这种格式，你可能需要实现一个自定义的http.RoundTripper。

// 但实际上，由于安全原因和标准化问题，许多代理服务器和库并不支持直接从URL中解析认证信息。
// 因此，你可能需要查找支持认证的代理客户端库，或者自己实现一个自定义的http.RoundTripper。

// 如果你选择自己实现，你需要在发送请求之前检查请求的目标是否是代理，
// 如果是，则添加一个带有认证信息的`Proxy-Authorization`头部。
// 这通常涉及到对`http.RoundTripper`接口的深入理解和实现。
