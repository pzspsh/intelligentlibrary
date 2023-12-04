/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:47:45
*/
package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Student struct {
	Name   string
	School string
}

type RpcServer struct{}

func (r *RpcServer) Introduce(student Student, words *string) error {
	fmt.Println("student: ", student)
	*words = fmt.Sprintf("Hello everyone, my name is %s, and I am from %s", student.Name, student.School)
	return nil
}

func main() {
	rpcServer := new(RpcServer)
	// 注册rpc服务
	_ = rpc.Register(rpcServer)
	//jsonrpc是基于TCP协议的，现在他还不支持http协议
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8082")
	tcpListen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	log.Println("tcp json-rpc service start success addr:8082")
	for {
		// 监听客户端请求
		conn, err3 := tcpListen.Accept()
		if err3 != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
