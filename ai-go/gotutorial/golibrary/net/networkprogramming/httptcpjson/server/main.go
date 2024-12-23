/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 17:06:56
*/
package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
)

// go对RPC的支持，支持三个级别：TCP、HTTP、JSONRPC
// go的RPC只支持GO开发的服务器与客户端之间的交互，因为采用了gob编码
// 注意字段必须是导出
type Params struct {
	Width, Height int
}
type Rect struct{}

// 函数必须是导出的
// 必须有两个导出类型参数
// 第一个参数是接收参数
// 第二个参数是返回给客户端参数，必须是指针类型
// 函数还要有一个返回值error
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}
func main() {
	rect := new(Rect)
	//注册一个rect服务
	rpc.Register(rect)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		//把服务处理绑定到http协议上
		rpc.HandleHTTP()
		err := http.ListenAndServe(":8080", nil)
		wg.Wait()
		if err != nil {
			log.Fatal(err)
			defer wg.Done()
		}
	}()
	log.Println("http rpc service start success addr:8080")

	go func() {
		tcpaddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
		tcplisten, err := net.ListenTCP("tcp", tcpaddr)
		if err != nil {
			log.Fatal(err)
			defer wg.Done()
		}
		for {
			conn, err3 := tcplisten.Accept()
			if err3 != nil {
				continue
			}
			go rpc.ServeConn(conn)
		}
	}()
	log.Println("tcp rpc service start success addr:8081")

	go func() {
		tcpaddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8082")
		tcplisten, err := net.ListenTCP("tcp", tcpaddr)
		if err != nil {
			log.Fatal(err)
			defer wg.Done()
		}
		for {
			conn, err3 := tcplisten.Accept()
			if err3 != nil {
				continue
			}
			go jsonrpc.ServeConn(conn)
		}
	}()
	log.Println("tcp json-rpc service start success addr:8082")
	wg.Wait()
}
