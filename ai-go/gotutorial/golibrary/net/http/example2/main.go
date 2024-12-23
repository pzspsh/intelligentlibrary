/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:44:19
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

// Http的三步骤
// 1、处理器函数
// 2、绑定URL关系
// 3、启动Server服务
func main() {
	//处理器函数
	timeFunc := func(response http.ResponseWriter, request *http.Request) {
		fmt.Println(request)
		now := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprint(response, now)
	}
	//绑定URL关系
	http.HandleFunc("/home/", timeFunc) //第一个参数为路径，第二个参数为匿名函数
	//启动web服务
	http.ListenAndServe(":9999", nil)
}
