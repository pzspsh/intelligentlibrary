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
	log.SetPrefix("client")
	file, err := os.OpenFile("Client.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	if err != nil {
		log.Println("Dial", err)
	}
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入你要发送的消息：")
	scan.Scan() //扫描
	fmt.Fprint(conn, scan.Text())
}
