/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 12:35:34
*/
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"time"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true, //这里是跳过证书验证，因为证书签发机构的CA证书是不被认证的
	}
	//注意这里要使用证书中包含的主机名称
	conn, err := tls.Dial("tcp", "127.0.0.1:8888", conf)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()
	log.Println("Client Connect To ", conn.RemoteAddr())
	status := conn.ConnectionState()
	fmt.Printf("%#v\n", status)
	buf := make([]byte, 1024)
	ticker := time.NewTicker(1 * time.Millisecond * 500)
	for {
		select {
		case <-ticker.C:
			{
				_, err = io.WriteString(conn, "hello")
				if err != nil {
					log.Fatalln(err.Error())
				}
				len, err := conn.Read(buf)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Receive From Server:", string(buf[:len]))
				}
			}
		}
	}
}
