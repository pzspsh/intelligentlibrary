/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 02:28:10
*/
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handle(c net.Conn) {
	// Handle the reads
	start := time.Now()
	tbuf := make([]byte, 81920)
	totalBytes := 0

	for {
		n, err := c.Read(tbuf)
		totalBytes += n
		// Was there an error in reading ?
		if err != nil {
			if err != io.EOF {
				log.Printf("Read error: %s", err)
			}
			break
		}
		log.Println(n)
	}
	log.Printf("%d bytes read in %s", totalBytes, time.Now().Sub(start))
	c.Close()
}

func main() {
	srv, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on localhost:2000")
	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}
