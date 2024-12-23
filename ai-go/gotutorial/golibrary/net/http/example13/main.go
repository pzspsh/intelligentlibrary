/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:44:19
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// 发送相关请求
func main() {
	//通过Get发送请求
	response, err := http.Get("http://localhost:9999/?a=1&b=2")
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(response.Proto, response.StatusCode, response.Status)
		//fmt.Println(response.Header)
		io.Copy(os.Stdout, response.Body)
	}
	//通过post发送请求
	buf := bytes.NewBufferString(`{"name":12}`) //创建一个缓冲对象
	resp, err := http.Post("localhost:9999", "application/json", buf)
	fmt.Println(resp, err)
	//通过PostForm请求
	parms := url.Values{}    //创建map映射
	parms.Add("name", "tom") //添加映射
	resp, err = http.PostForm("http://localhost:8888", parms)
	fmt.Println(resp, err)
}
