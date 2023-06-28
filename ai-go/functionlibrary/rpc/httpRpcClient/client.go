/*
@File   : client.go
@Author : pan
@Time   : 2023-06-26 14:17:58
*/
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A int
	B int
}

type Quotient struct {
	Quo int
	Rem int
}

// \rpc\client> go run client.go localhost
func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage:", os.Args[0], "server")
		os.Exit(1)
	}

	serverAddress := os.Args[1]
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Arith:%d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		fmt.Printf("client call err:%v", err)
	}
	fmt.Printf("Arith:%d/%d=%d......%d\n", args.A, args.B, quot.Quo, quot.Rem)

}
