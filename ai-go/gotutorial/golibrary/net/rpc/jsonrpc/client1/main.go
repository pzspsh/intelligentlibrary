/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:47:45
*/
package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main() {
	type Student struct {
		Name   string
		School string
	}
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:8082")
	if err != nil {
		panic(err)
	}
	defer client.Close()
	var reply string
	// 发送json格式的数据
	err = client.Call("RpcServer.Introduce", &Student{
		Name:   "random_w",
		School: "Secret",
	}, &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
