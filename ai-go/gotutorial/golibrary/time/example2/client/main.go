/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 02:12:49
*/
package main

import (
	"log"
	"net"
	"os"
	"time"
)

func main() {
	connTimeout := 3 * time.Second
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8080", connTimeout) // 3s timeout
	if err != nil {
		log.Println("dial failed:", err)
		os.Exit(1)
	}
	defer conn.Close()
	readTimeout := 2 * time.Second
	buffer := make([]byte, 512)
	for {
		err = conn.SetReadDeadline(time.Now().Add(readTimeout)) // timeout
		if err != nil {
			log.Println("setReadDeadline failed:", err)
		}
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Read failed:", err)
			//break
		}
		log.Println("count:", n, "msg:", string(buffer))
	}
}
