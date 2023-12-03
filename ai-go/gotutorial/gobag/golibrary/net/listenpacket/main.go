/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:13:21
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	pc, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go handlePacket(pc, addr, buf[:n])
	}
}

func handlePacket(pc net.PacketConn, addr net.Addr, buf []byte) {
	//TODO: process packet
}
