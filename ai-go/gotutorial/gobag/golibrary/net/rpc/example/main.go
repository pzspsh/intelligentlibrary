/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:58:17
*/
package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Greeter struct{}

func (g *Greeter) Greet(name *string, reply *string) error {
	*reply = fmt.Sprintf("Hello, %s!", *name)
	return nil
}

func main() {
	addr := "localhost:12345"

	greeter := new(Greeter)
	rpc.Register(greeter)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("RPC server listening on %s\n", addr)

	rpc.Accept(listener)
}
