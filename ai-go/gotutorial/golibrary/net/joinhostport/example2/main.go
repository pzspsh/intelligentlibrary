/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 14:11:22
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.InterfaceAddrs()
	fmt.Println(addr) //[127.0.0.1/8 10.236.15.24/24 ::1/128 fe80::3617:ebff:febe:f123/64],本地地址,ipv4和ipv6地址,这些信息可通过ifconfig命令看到
	interfaces, _ := net.Interfaces()
	fmt.Println(interfaces) //[{1 65536 lo  up|loopBACk} {2 1500 eth0 34:17:eb:be:f1:23 up|broadcast|multicast}] 类型:MTU(最大传输单元),网络接口名,支持状态
	hp := net.JoinHostPort("127.0.0.1", "8080")
	fmt.Println(hp) //127.0.0.1:8080,根据ip和端口组成一个addr字符串表示
	lt, _ := net.LookupAddr("127.0.0.1")
	fmt.Println(lt) //[localhost],根据地址查找到改地址的一个映射列表
	cname, _ := net.LookupCNAME("www.baidu.com")
	fmt.Println(cname) //www.a.shifen.com,查找规范的dns主机名字
	host, _ := net.LookupHost("www.baidu.com")
	fmt.Println(host) //[111.13.100.92 111.13.100.91],查找给定域名的host名称
	ip, _ := net.LookupIP("www.baidu.com")
	fmt.Println(ip) //[111.13.100.92 111.13.100.91],查找给定域名的ip地址,可通过nslookup www.baidu.com进行查找操作.
}
