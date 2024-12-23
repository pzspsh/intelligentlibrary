/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 14:49:47
*/
package main

import (
	"fmt"
	"os/exec"
)

/*
使用http代理的动态数据请求
*/
func Third(webUrl, jsFileName, proxyHost, proxyPort string) {
	// cmd := exec.Command("phantomjs.exe", jsFileName, webUrl)
	cmd := exec.Command("phantomjs.exe", jsFileName, proxyHost, proxyPort, webUrl)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(out))
}
func main() {
	webUrl := "http://httpbin.org/ip" //"http://httpbin.org/ip" // "http://ip.gs/" // "http://httpbin.org/user-agent"
	jsfileName := "somescript.js"
	Third(webUrl, jsfileName, "139.224.237.33", "8888")
}
