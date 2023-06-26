/*
@File   : server.go
@Author : pan
@Time   : 2023-06-26 14:17:39
*/
package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

// 参数
type Args struct {
	A, B int
}

// 结果
type Quotient struct {
	Quo int //求商
	Rem int //求余数
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	return nil
}

func main() {
	arith := new(Arith)

	rpc.Register(arith)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
