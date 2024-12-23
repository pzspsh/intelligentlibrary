/*
@File   : server.go
@Author : pan
@Time   : 2023-06-29 14:44:16
*/
package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// rpc服务端
// 定义一个远程调用的结构体,并创建一个远程调用的函数,函数一般是放在结构体中的
type Hello struct {
}

/*
说明:

	1、方法只能有两个可序列化的参数，其中第二个参数是指针类型
	    req 表示获取客户端传过来的数据
	    res 表示给客户端返回数据
	2、方法要返回一个error类型，同时必须是公开的方法
	3、req和res的类型不能是：channel（通道）、func（函数）,因为以上类型均不能进行 序列化
*/
func (h Hello) SayHello(req string, res *string) error {
	fmt.Println("请求的参数:", req)
	// 设置返回的数据
	*res = "你好" + req
	return nil
}

func main() {
	// 1、注册RPC服务
	// hello: rpc服务名称
	err1 := rpc.RegisterName("hello", new(Hello))
	if err1 != nil {
		fmt.Println(err1)
	}

	// 2、监听端口
	listen, err2 := net.Listen("tcp", "127.0.0.1:8080")
	if err2 != nil {
		fmt.Println(err2)
	}

	// 3、应用退出的时候关闭监听端口
	defer listen.Close()

	for {
		// for 循环, 一直进行连接,每个客户端都可以连接
		fmt.Println("开始创建连接")
		// 4、建立连接
		conn, err3 := listen.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}
		// 5、绑定服务
		rpc.ServeConn(conn)
	}
}
