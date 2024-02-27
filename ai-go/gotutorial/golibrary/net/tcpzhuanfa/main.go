/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 00:54:17
*/
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handler(src net.Conn) {
	defer src.Close()
	dst, err := net.DialTimeout("tcp", "192.168.135.128:8080", time.Second*6)
	if err != nil {
		log.Fatalln("Unable to connect to out host")
	}
	defer dst.Close()
	log.Printf("real addr:%s\n", dst.RemoteAddr())
	if err := Transport(src, dst); err != nil {
		log.Fatalf("exchange err:%s", err)
	}
	log.Printf("data exchange over")

}

func Transport(rw1, rw2 io.ReadWriter) error {
	errc := make(chan error, 1)
	go func() {
		errc <- copyBuffer(rw1, rw2)
	}()

	go func() {
		errc <- copyBuffer(rw2, rw1)
	}()

	err := <-errc
	if err != nil && err == io.EOF {
		err = nil
	}
	return err
}

func copyBuffer(dst io.Writer, src io.Reader) error {
	_, err := io.Copy(dst, src)
	return err
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("get a conn")
		go handler(conn)
	}
}
