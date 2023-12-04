/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:47:45
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type MathService struct {
}

type Args struct {
	A, B int
}

func (m *MathService) Add(args Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func main() {
	client, err := DialHTTP("tcp", "localhost:8088")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Args{A: 1, B: 2}
	var reply int
	err = client.Call("MathService.Add", args, &reply)
	if err != nil {
		log.Fatal("MathService.Add error:", err)
	}
	fmt.Printf("MathService.Add: %d+%d=%d", args.A, args.B, reply)
}

// DialHTTP connects to an HTTP RPC server at the specified network address
// listening on the default HTTP RPC path.
func DialHTTP(network, address string) (*rpc.Client, error) {
	return DialHTTPPath(network, address, rpc.DefaultRPCPath)
}

// DialHTTPPath connects to an HTTP RPC server
// at the specified network address and path.
func DialHTTPPath(network, address, path string) (*rpc.Client, error) {
	var err error
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	io.WriteString(conn, "GET "+path+" HTTP/1.0\n\n")
	// Require successful HTTP response
	// before switching to RPC protocol.
	resp, err := http.ReadResponse(bufio.NewReader(conn), &http.Request{Method: "GET"})
	connected := "200 Connected to JSON RPC"
	if err == nil && resp.Status == connected {
		return jsonrpc.NewClient(conn), nil
	}
	if err == nil {
		err = errors.New("unexpected HTTP response: " + resp.Status)
	}
	conn.Close()
	return nil, &net.OpError{
		Op:   "dial-http",
		Net:  network + " " + address,
		Addr: nil,
		Err:  err,
	}
}
