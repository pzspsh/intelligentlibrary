package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	type Student struct {
		Name   string
		School string
	}
	// 连接RPC服务端 Dial会调用NewClient初始化一个Client
	client, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	defer client.Close()
	// 发送请求
	var reply string
	err = client.Call("RpcServer.Introduce", &Student{Name: "random_w", School: "Secret"}, &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
