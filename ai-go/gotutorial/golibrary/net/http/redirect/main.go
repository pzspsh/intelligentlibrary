/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:15:57
*/
package main

import (
	"fmt"
	"net/http"
)

func RedirectReq() {
	client := &http.Client{}
	url := "http://www.qq.com"
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	response, _ := client.Do(reqest)
	fmt.Println(response.StatusCode)
}

func DemoRedirect(enableRedirect bool) {
	client := &http.Client{}

	if !enableRedirect {
		// 关闭重定向
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	} else {
		// 启用重定向
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			fmt.Printf("重定向到: %s\n", req.URL)
			return nil
		}
	}

	// 发起请求
	resp, err := client.Get("http://httpbin.org/redirect/1") // 测试用的重定向 URL
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 输出响应状态码和最终 URL
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("最终 URL: %s\n", resp.Request.URL)
}

func RedirectRun() {
	// 测试关闭重定向
	fmt.Println("=== 测试关闭重定向 ===")
	DemoRedirect(false)

	// 测试启用重定向
	fmt.Println("\n=== 测试启用重定向 ===")
	DemoRedirect(true)

	/*
	   === 测试关闭重定向 ===
	   状态码: 302
	   最终 URL: http://httpbin.org/redirect/1


	   === 测试启用重定向 ===
	   重定向到: http://httpbin.org/get
	   状态码: 200
	   最终 URL: http://httpbin.org/get
	*/
}

func StopRedirect() { // 关闭重定向
	/*
		通过设置 CheckRedirect 函数返回一个错误（如 http.ErrUseLastResponse），可以阻止 http.Client 自动跟随重定向。
	*/
	// 创建一个自定义的 HTTP 客户端，关闭重定向
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 返回 http.ErrUseLastResponse 表示不跟随重定向
			return http.ErrUseLastResponse
		},
	}

	// 发起请求
	resp, err := client.Get("http://example.com")
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 输出响应状态码和 URL
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("最终 URL: %s\n", resp.Request.URL)
}

func EnableRedirect() { // 启用重定向
	/*
		默认情况下，http.Client 会自动处理重定向。你无需额外配置即可启用重定向。如果需要自定义重定向行为，可以在 CheckRedirect 中实现逻辑。
	*/
	// 创建一个自定义的 HTTP 客户端，默认启用重定向
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 打印每次重定向的信息
			fmt.Printf("重定向到: %s\n", req.URL)
			return nil // 返回 nil 表示允许重定向
		},
	}

	// 发起请求
	resp, err := client.Get("http://example.com")
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 输出响应状态码和最终 URL
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("最终 URL: %s\n", resp.Request.URL)
}

func main() {
	// RedirectReq()
	// EnableRedirect()
	RedirectRun()
}

/*
关键点解析
	CheckRedirect 函数的作用：
		CheckRedirect 是 http.Client 的一个字段，用于控制如何处理重定向。
		如果返回 nil，表示允许重定向。
		如果返回错误（如 http.ErrUseLastResponse），表示禁止重定向。
	默认行为：
		如果未设置 CheckRedirect，http.Client 会自动处理最多 10 次重定向（超过 10 次会报错）。
	http.ErrUseLastResponse：
		这是一个特殊的错误值，表示不跟随重定向，直接返回原始响应。
	调试重定向：
		在 CheckRedirect 中可以打印每次重定向的 URL 和其他信息，便于调试。
*/
