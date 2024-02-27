/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 02:11:09
*/
package main

import (
	"log"
	"net"
	"time"
)

func main() {
	addr := "0.0.0.0:8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalf("net.ResovleTCPAddr fail:%s", addr)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("listen %s fail: %s", addr, err)
	} else {
		log.Println("listening", addr)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listener.Accept error:", err)
			continue
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	var buffer []byte = []byte("You are welcome. I'm server.")
	for {
		time.Sleep(3 * time.Second) // sleep 3s
		n, err := conn.Write(buffer)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
		log.Println("send:", n)
	}
	log.Println("connetion end")
}
