/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 19:53:11
*/
package main

import (
	"fmt"
	"net"
)

func main() {

	// udp server
	listenUdp, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("ListenPacket fail err", err)
		return
	}
	defer listenUdp.Close()

	for {
		// 接收数据
		var buf [1024]byte
		n, addr, err := listenUdp.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("ReadFromUDP fail err", err)
			continue
		}
		fmt.Printf("接收到客户端【%v】的数据：%s\n", addr, string(buf[:n]))

		// 发送数据
		msg := "server发送回来的数据" + string(buf[:n])
		num, err := listenUdp.WriteToUDP([]byte(msg), addr)
		if err != nil {
			fmt.Println("WriteToUDP fail err", num, err)
			continue
		}
	}
}
