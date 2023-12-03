/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:58:50
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func connectControl() {
	var tcpAddr *net.TCPAddr
	//这里在一台机测试，所以没有连接到公网，可以修改到公网ip
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:8009")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return
	}
	fmt.Println(conn.LocalAddr().String() + " : Client connected!8009")
	reader := bufio.NewReader(conn)
	for {
		s, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		} else {
			//接收到new的指令的时候，新建一个tcp连接
			if s == "new\n" {
				go combine()
			}
			if s == "hi" {
				//忽略掉hi的请求
			}
		}

	}
}

func combine() {
	local := connectLocal()
	remote := connectRemote()
	if local != nil && remote != nil {
		joinConn(local, remote)
	} else {
		if local != nil {
			err := local.Close()
			if err != nil {
				fmt.Println("close local:" + err.Error())
			}
		}
		if remote != nil {
			err := remote.Close()
			if err != nil {
				fmt.Println("close remote:" + err.Error())
			}

		}
	}
}
func joinConn(local *net.TCPConn, remote *net.TCPConn) {
	f := func(local *net.TCPConn, remote *net.TCPConn) {
		defer local.Close()
		defer remote.Close()
		_, err := io.Copy(local, remote)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("end")
	}
	go f(local, remote)
	go f(remote, local)
}

func connectLocal() *net.TCPConn {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:8000")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return nil
	}

	fmt.Println(conn.LocalAddr().String() + " : Client connected!8000")
	return conn

}

func connectRemote() *net.TCPConn {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:8008")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return nil
	}
	fmt.Println(conn.LocalAddr().String() + " : Client connected!8008")
	return conn
}

func main() {
	connectControl()
}
