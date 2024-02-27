/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:42:34
*/
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	file, err := os.Create("ServerLog.log")
	log.SetPrefix("server:")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	addr := "127.0.0.1:8080"
	packetConn, err := net.ListenPacket("udp", addr) //使用UDP的方式监听
	if err != nil {
		fmt.Println(err)
	}
	defer packetConn.Close()
	if err != nil {
		log.Println(err)
	}
	ctx := make([]byte, 1024)
	for {
		n, addr, err := packetConn.ReadFrom(ctx)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(addr.String()) //传输过来的地址
		fmt.Println(string(ctx[0:n]))
		packetConn.WriteTo([]byte("测试UDP传输回去"), addr)
	}
}
