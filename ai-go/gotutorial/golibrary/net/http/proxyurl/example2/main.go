/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 14:47:34
*/
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func First() {
	/*
		1. 普通请求
	*/

	webUrl := "http://ip.gs/"
	resp, err := http.Get(webUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	// if resp.StatusCode == http.StatusOK {
	// 	fmt.Println(resp.StatusCode)
	// }

	time.Sleep(time.Second * 3)

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func Second(webUrl, proxyUrl string) {
	/*
		1. 代理请求
		2. 跳过https不安全验证
	*/
	// webUrl := "http://ip.gs/"
	// proxyUrl := "http://115.215.71.12:808"

	proxy, _ := url.Parse(proxyUrl)
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5, //超时时间
	}

	resp, err := client.Get(webUrl)
	if err != nil {
		fmt.Println("出错了", err)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func Third(webUrl, proxyUrl string) {
	/*
		1. 代理请求
		2. 跳过https不安全验证
		3. 自定义请求头 User-Agent

	*/
	// webUrl := "http://ip.gs/"
	// proxyUrl := "http://171.215.227.125:9000"

	request, _ := http.NewRequest("GET", webUrl, nil)
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")

	proxy, _ := url.Parse(proxyUrl)
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5, //超时时间
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("出错了", err)
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func main() {
	webUrl := "http://httpbin.org/user-agent" //"http://ip.gs/"
	proxyUrl := "http://119.5.0.75:808"

	Second(webUrl, proxyUrl)
	// Third(webUrl, proxyUrl)
}
