/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:47:55
*/
package main

import (
	"fmt"
	"log"
	"net"
)

func chkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//创建一个TCP服务端
	ta, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8081")
	chkError(err)
	//监听端口
	tl, err2 := net.ListenTCP("tcp", ta)
	chkError(err2)
	fmt.Println("Server Start")
	//.建立链接并处理
	go func() {
		for {
			//如果有客户端链接过来，阻塞会返回
			conn, err := tl.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			//已经与客户端建立链接，处理业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}
					//回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write bak buf err", err)
						continue
					}
				}
			}()
		}
	}()
	//阻塞状态
	select {}
}
