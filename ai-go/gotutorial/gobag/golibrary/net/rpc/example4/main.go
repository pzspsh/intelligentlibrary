/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:00:41
*/
package main

import (
	"fmt"
	"net"
	"net/rpc"

	"net/rpc/jsonrpc"
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
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
