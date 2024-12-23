/*
@File   : client.go
@Author : pan
@Time   : 2023-06-28 16:32:20
*/
package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 算数运算响应结构体
type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

func main() {
	conn, err := jsonrpc.Dial("tcp", "39.105.119.219:8096")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	req := ArithRequest{9, 2}
	var res ArithResponse

	err = conn.Call("Arith.Multiply", req, &res) // 乘法运算
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d * %d = %d \n", req.A, req.B, res.Pro)

	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %dn", req.A, req.B, res.Quo, res.Rem)
}
