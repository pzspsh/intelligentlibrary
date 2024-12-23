/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:42:10
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	rs, err := http.Get("https://baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	//读取响应头信息
	fmt.Println("rs.Status:", rs.Status)
	fmt.Println("rs.StatusCode:", rs.StatusCode)
	fmt.Println("rs.TransferEncoding:", rs.TransferEncoding)
	fmt.Println("rs.Proto:", rs.Proto)
	fmt.Println("rs.ProtoMajor:", rs.ProtoMajor)
	fmt.Println("rs.ProtoMinor:", rs.ProtoMinor)
	fmt.Println("rs.Request:", rs.Request)
	fmt.Println("rs.ContentLength:", rs.ContentLength)
	fmt.Println("rs.Header:", rs.Header)
	fmt.Println("rs.Close:", rs.Close)
	//读取body中的数据
	fmt.Println()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer rs.Body.Close()
	fmt.Println(string(body))
}
