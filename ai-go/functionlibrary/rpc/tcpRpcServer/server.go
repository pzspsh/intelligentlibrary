/*
@File   : server.go
@Author : pan
@Time   : 2023-06-29 14:15:58
*/
package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义类对象
type World struct{}

// 绑定类方法
func (w *World) HelloWorld(req string, res *string) error {
	*res = req + " 你好!"
	return nil
}

func main() {
	// 1. 注册RPC服务
	err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println("注册 rpc 服务失败!", err)
		return
	}
	// 2. 设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	defer listener.Close()
	fmt.Println("开始监听 ...")

	// 3. 建立链接
	for {
		//接收连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept() err:", err)
			return
		}
		// 4. 绑定服务
		go rpc.ServeConn(conn)
	}
}
