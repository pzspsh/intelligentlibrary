/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:54:36
*/
package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
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
	// 指定rpc模式为TCP模式，地址为127.0.0.1:8081
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	tcpListen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("tcp rpc service start success addr:8081")
	for {
		// 监听Client发送的请求
		conn, err3 := tcpListen.Accept()
		if err3 != nil {
			continue
		}
		// 创建一个goroutine处理请求
		go rpc.ServeConn(conn)
	}
}
