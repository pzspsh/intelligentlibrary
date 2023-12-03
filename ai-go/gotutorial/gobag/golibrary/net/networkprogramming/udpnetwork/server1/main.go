/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 02:05:56
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

const (
	send string = "send the data"
	isOk string = "is ok"
)

func main() {
	var port string
	flag.StringVar(&port, "addr", ":6666", "Address is listen to")
	fmt.Println("Listener to 127.0.0.1:6666")
	lp, err := net.ListenPacket("udp", port)
	if err != nil {
		log.Fatalf("connect is error %v", err)
	}
	defer lp.Close()
	for {
		buf := make([]byte, 1024) // 固定值可以避免粘包问题
		n, addr, err := lp.ReadFrom(buf)
		if err != nil {
			continue
		}
		fmt.Println("服务端接收的信息: ", string(buf[:n]))
		var send = append([]byte("<----- udp < 服务器已经接收到您的消息："), string(buf[:n])...)
		go server(lp, send, addr)
	}
}

func server(lp net.PacketConn, buf []byte, addr net.Addr) {
	buf[2] |= 0x1 // 操作第3个字节从P变q
	// 此处可以做很多逻辑来保持udp长连接，但是此处直接返回信息
	lp.WriteTo(buf, addr)
}
