/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 19:53:11
*/
package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// udp client
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("DialUDP fail err", err)
		return
	}

	defer udpConn.Close()

	for {
		// 发送数据
		var data string
		fmt.Print("请输入要发送给服务器的数据：")
		fmt.Scanln(&data)
		if strings.ToUpper(data) == "Q" {
			break
		}
		sendData := []byte(data)
		_, err = udpConn.Write(sendData)
		if err != nil {
			fmt.Println("udp客户端Write fail err", err)
			return
		}

		// 接收数据
		buf := make([]byte, 4096)
		n, remoteAddr, err2 := udpConn.ReadFromUDP(buf[:])
		if err2 != nil {
			fmt.Println("udp客户端Read fail err", err2)
			return
		}
		fmt.Printf("data：%v, addr：%v\n", string(buf[:n]), remoteAddr)

	}
}
