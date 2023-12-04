/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 10:04:36
*/
package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
)

/*
生成私钥：openssl genrsa -out key.pem 2048
生成证书：openssl req -new -x509 -key key.pem -out cert.pem -days 3650
*/

func main() {
	log.SetFlags(log.Lshortfile)
	cer, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Println(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", ":8000", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		println(msg)
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
