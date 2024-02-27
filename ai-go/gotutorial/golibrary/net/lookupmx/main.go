/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 14:23:21
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	// 查找地址映射列表

	//主机名的dns  的mx 记录 邮件服务器地址
	mx, error := net.LookupMX("baidu.com")
	fmt.Println(error)
	for _, value := range mx {
		fmt.Println(value.Host)
	}
	//主机名的dns  的ns 记录 解析该域名的服务器地址
	ns, error := net.LookupNS("mingbozhu.com")
	fmt.Println(error)
	fmt.Println(ns[0].Host)

	// 联系人信息
	txt, error := net.LookupTXT("mingbozhu.com")
	fmt.Println(error)
	fmt.Println(txt)
}
