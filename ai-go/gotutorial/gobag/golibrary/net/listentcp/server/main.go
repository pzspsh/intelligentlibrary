/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:38:55
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	log.SetPrefix("main")
	file, err := os.OpenFile("Server.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	if err != nil {
		log.Fatal("OpenFile:", err)
	}
	log.SetOutput(file)
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Println(err)
	}
	conn, err := listen.Accept()
	if err != nil {
		log.Println("Accept", err)
	}
	defer conn.Close()
	for {
		data := make([]byte, 1024)
		reader := bufio.NewReader(conn)
		n, err := reader.Read(data)
		if err != nil {
			log.Println("buf", err)
			break
		}
		fmt.Println(string(data[0:n]))
	}
}
