/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:58:27
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
	addr := "127.0.0.1:1024"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	log.Printf("connection server: %s success", addr)
	sendLoop(conn)
}

func send(conn net.Conn) {
	words := "hello server!"
	conn.Write([]byte(words))
	log.Println("send over")

	// receive from server
	buffer := make([]byte, 2048)

	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("%s waiting server back msg error: %s", conn.RemoteAddr(), err)
		return
	}
	log.Printf("%s receive server back msg: %s", conn.RemoteAddr(), string(buffer[:n]))
}

func sendLoop(conn net.Conn) {
	for {
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		_, err := conn.Write([]byte(input))
		if err != nil {
			log.Println(err)
		}
		log.Println("send over")

		// receive from server
		buffer := make([]byte, 2048)

		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("%s waiting server back msg error: %s", conn.RemoteAddr(), err)
			return
		}
		log.Printf("%s receive server back msg: %s", conn.RemoteAddr(), string(buffer[:n]))
	}

}
