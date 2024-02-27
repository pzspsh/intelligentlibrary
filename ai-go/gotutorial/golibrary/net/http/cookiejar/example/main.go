/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:08:38
*/
package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func main() {
	// 创建cookiejar实例
	cookieJar, _ := cookiejar.New(nil)

	// 创建http.Client实例，并设置cookiejar
	httpClient := &http.Client{
		Jar: cookieJar,
	}

	// 发送http请求
	resp, err := httpClient.Get("http://example.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Body.Close()

	// 在之后的http请求中，会自动使用cookiejar中的cookie
	resp2, err := httpClient.Get("http://example.com/profile")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp2.Body.Close()

	// 手动添加cookie
	url, _ := url.Parse("http://example.com")
	cookie := &http.Cookie{Name: "myCookie", Value: "myValue"}
	cookieJar.SetCookies(url, []*http.Cookie{cookie})

	// 获取所有cookie
	cookies := cookieJar.Cookies(url)
	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s\n", cookie.Name, cookie.Value)
	}
}
