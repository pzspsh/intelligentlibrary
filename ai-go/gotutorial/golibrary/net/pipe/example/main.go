/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 14:47:27
*/
package main

import (
	"bytes"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

func echoHandle(conn net.Conn) {
	defer conn.Close()
	buf := bytes.Buffer{}

	for {
		piece := make([]byte, 256)

		n, err := conn.Read(piece)

		piece = piece[:n]

		if err != nil {
			break
		}

		buf.Write(piece)

		n, err = conn.Write(piece)
		if err != nil {
			break
		}

		buf.Next(n)

		log.WithField("count", n).Trace("echo several bytes to client")
	}

	log.Debug("connection closed.")
}

func main() {
	exitCh := make(chan interface{})
	cli, serv := net.Pipe()
	go echoHandle(serv)

	kText := "hello"
	buf := bytes.Buffer{}
	piece := make([]byte, 256)
	buf.WriteString(kText)

	for buf.Len() != 0 {
		n, err := cli.Write(buf.Bytes())
		if err != nil {
			log.Debug(err)
			break
		}
		buf.Next(n)
	}

	n, err := cli.Read(piece)
	if err != nil {
		log.Debug(err)
	}
	piece = piece[:n]
	close(exitCh)

	if string(piece) != kText {
		fmt.Printf("got a broken string from echo service. expected=%v actual=%v\n", kText, buf.String())
	}
}
