/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 16:33:10
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	ip, _ := net.LookupIP("www.baidu.com")
	fmt.Println(ip) //[111.13.100.92 111.13.100.91],查找给定域名的ip地址,可通过nslookup www.baidu.com进行查找操作.
}
