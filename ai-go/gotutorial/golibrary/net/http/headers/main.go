/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 10:39:40
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	head := http.Header{}
	fmt.Println(head)
	//添加
	head.Add("User-Agent", "Chrome")
	head.Add("Content-Type", "[text/html]")
	fmt.Println(head)
	//获取
	fmt.Println(head.Get("User-Agent"))
	//设置
	head.Set("User-Agent", "QQ Browser")
	fmt.Println(head)
	//删除
	head.Del("User-Agent")
	fmt.Println("head.Values", head.Values("User-Agent"))
	fmt.Println(head)
}
