/*
@File   : server.go
@Author : pan
@Time   : 2023-11-09 11:10:23
*/
package socket5proxy

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	var b [1024]byte

	n, err := client.Read(b[:])
	if err != nil {
		log.Printf("n:%d ,Client Read Error:%v", n, err)
		return
	}

	if b[0] == 0x05 { //只处理Socks5协议
		//客户端回应：Socks服务端不需要验证方式
		client.Write([]byte{0x05, 0x00})
		n, err = client.Read(b[:])
		if err != nil {
			log.Printf("client Read Error:%v", err)
			return
		}
		var host, port string
		switch b[3] {
		case 0x01: //IP V4
			host = net.IPv4(b[4], b[5], b[6], b[7]).String()
		case 0x03: //域名
			host = string(b[5 : n-2]) //b[4]表示域名的长度
		case 0x04: //IP V6
			host = net.IP{b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19]}.String()
		}
		port = strconv.Itoa(int(b[n-2])<<8 | int(b[n-1]))
		server, err := net.Dial("tcp", net.JoinHostPort(host, port))
		fmt.Println("server", server, host, port)
		if err != nil {
			log.Println(err)
			return
		}
		defer server.Close()
		client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) //响应客户端连接成功
		//进行转发
		go io.Copy(server, client)
		io.Copy(client, server)
	}
}

func StartProxyServer(proxyip, proxyport string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", proxyip, proxyport))
	if err != nil {
		log.Panic(err)
	}
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(client)
	}
}
