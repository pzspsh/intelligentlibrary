/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:42:34
*/
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var str string
	file, err := os.Create("ClientErr.log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		log.Println(err)
	}
	fmt.Scanln(&str)
	fmt.Fprintf(conn, "%s", str)
	ctx := make([]byte, 1024)
	n, err := conn.Read(ctx)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(ctx[0:n]))
}
