/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:47:45
*/
package main

import (
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
	rpc.RegisterName("MathService", new(MathService))
	//注册一个path，用于提供基于http的json rpc服务
	http.HandleFunc(rpc.DefaultRPCPath, func(rw http.ResponseWriter, r *http.Request) {
		conn, _, err := rw.(http.Hijacker).Hijack()
		if err != nil {
			log.Print("rpc hijacking ", r.RemoteAddr, ": ", err.Error())
			return
		}
		var connected = "200 Connected to JSON RPC"
		io.WriteString(conn, "HTTP/1.0 "+connected+"\n\n")
		jsonrpc.ServeConn(conn)
	})
	l, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil) //换成http的服务
}
