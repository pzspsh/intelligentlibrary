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

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	reply := ""
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c))
	err = client.Call("FoodService.SayName", "鱼香肉丝", &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply)
}
