/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 02:28:10
*/
package main

import (
	"log"
	"net"
	"time"
)

func handle(c net.Conn) {
	start := time.Now()
	tbuf := make([]byte, 4096)
	totalBytes := 0
	for i := 0; i < 1000; i++ {
		n, err := c.Write(tbuf)
		totalBytes += n
		// Was there an error in writing?
		if err != nil {
			log.Printf("Write error: %s", err)
			break
		}
		log.Println(n)
	}
	log.Printf("%d bytes written in %s", totalBytes, time.Now().Sub(start))
	c.Close()
}

func main() {
	conn, err := net.Dial("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Sending to localhost:2000")
	handle(conn)
}
