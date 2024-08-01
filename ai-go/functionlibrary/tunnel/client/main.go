/*
@File   : main.go
@Author : pan
@Time   : 2024-08-01 11:50:38
*/
package main

import (
	"bufio"
	"io"
	"log"
	"net"

	"function/tunnel/network"
)

var (
	// 本地需要暴露的服务端口
	localServerAddr = "127.0.0.1:32768"

	remoteIP = "111.111.111.111"
	// 远端的服务控制通道，用来传递控制信息，如出现新连接和心跳
	remoteControlAddr = remoteIP + ":8009"
	// 远端服务端口，用来建立隧道
	remoteServerAddr = remoteIP + ":8008"
)

func main() {
	tcpConn, err := network.CreateTCPConn(remoteControlAddr)
	if err != nil {
		log.Println("[连接失败]" + remoteControlAddr + err.Error())
		return
	}
	log.Println("[已连接]" + remoteControlAddr)

	reader := bufio.NewReader(tcpConn)
	for {
		s, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}

		// 当有新连接信号出现时，新建一个tcp连接
		if s == network.NewConnection+"\n" {
			go connectLocalAndRemote()
		}
	}

	log.Println("[已断开]" + remoteControlAddr)
}

func connectLocalAndRemote() {
	local := connectLocal()
	remote := connectRemote()

	if local != nil && remote != nil {
		network.Join2Conn(local, remote)
	} else {
		if local != nil {
			_ = local.Close()
		}
		if remote != nil {
			_ = remote.Close()
		}
	}
}

func connectLocal() *net.TCPConn {
	conn, err := network.CreateTCPConn(localServerAddr)
	if err != nil {
		log.Println("[连接本地服务失败]" + err.Error())
	}
	return conn
}

func connectRemote() *net.TCPConn {
	conn, err := network.CreateTCPConn(remoteServerAddr)
	if err != nil {
		log.Println("[连接远端服务失败]" + err.Error())
	}
	return conn
}
