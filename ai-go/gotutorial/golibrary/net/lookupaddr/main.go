package main

import (
	"fmt"
	"net"
)

func main() {
	lt, _ := net.LookupAddr("127.0.0.1")
	fmt.Println(lt) //[localhost],根据地址查找到改地址的一个映射列表
}
