/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:47:45
*/
package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type FoodService struct{}

func (f *FoodService) SayName(request string, resp *string) error {
	fmt.Println("开始执行方法")
	*resp = "你想吃：" + request
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = rpc.RegisterName("FoodService", &FoodService{})
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("处理客户端连接")
		// rpc.ServeConn(conn)
		// fmt.Println("连接关闭了")
		fmt.Println("开启go协程")
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
