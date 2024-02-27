/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 12:22:02
*/
package main

import (
	"crypto/tls"
	"log"
)

func main() {
	cert, err := tls.LoadX509KeyPair("testdata/example-cert.pem", "testdata/example-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":2000", cfg)
	if err != nil {
		log.Fatal(err)
	}
	_ = listener
}
