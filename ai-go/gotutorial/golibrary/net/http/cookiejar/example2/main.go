/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:13:03
*/
package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	// 创建cookiejar.Jar对象
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	// 创建http.Client对象，并设置其Jar属性为上面创建的cookiejar.Jar对象
	client := http.Client{
		Jar: jar,
	}

	// 发送GET请求
	resp, err := client.Get("https://www.example.com")
	if err != nil {
		panic(err)
	}

	// 输出响应内容
	fmt.Println(resp.Status)

	// 打印请求返回的所有Cookie
	cookies := jar.Cookies(resp.Request.URL)
	for _, cookie := range cookies {
		fmt.Printf("Cookie: %s=%s", cookie.Name, cookie.Value)
	}
}
