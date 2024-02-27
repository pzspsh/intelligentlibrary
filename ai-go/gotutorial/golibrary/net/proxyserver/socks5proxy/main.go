/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 15:40:41
*/
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/armon/go-socks5"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 58888, "Port for the proxy server")
	flag.Parse()

	// 创建 SOCKS5 服务器实例
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatalf("Failed to create SOCKS5 server: %v", err)
	}

	// 监听并处理请求
	addr := fmt.Sprintf(":%d", port)
	log.Printf("SOCKS5 proxy server started on port %d\n", port)
	if err := server.ListenAndServe("tcp", addr); err != nil {
		log.Fatalf("Failed to start SOCKS5 server: %v", err)
	}
}
