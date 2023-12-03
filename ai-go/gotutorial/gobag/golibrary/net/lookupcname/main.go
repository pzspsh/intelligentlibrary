/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 14:16:00
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	cname, _ := net.LookupCNAME("www.baidu.com")
	fmt.Println(cname) //www.a.shifen.com,查找规范的dns主机名字
}
