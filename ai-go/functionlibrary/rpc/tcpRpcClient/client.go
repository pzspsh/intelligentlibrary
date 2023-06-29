/*
@File   : client.go
@Author : pan
@Time   : 2023-06-29 14:15:36
*/
package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. 用 rpc 链接服务器 --Dial()
	conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}

	defer conn.Close()

	// 2. 调用远程函数
	var reply string // 接受返回值 --- 传出参数
	err = conn.Call("hello.Hello World", "张三", &reply)
	if err != nil {
		fmt.Println("Call:", err)
		return
	}
	fmt.Println(reply)
}
