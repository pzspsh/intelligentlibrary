package main

import (
	"fmt"
	"net"
	"net/http"
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
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("RPC server listening on %s\n", addr)
	http.Serve(listener, nil)
}
