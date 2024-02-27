/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 00:02:58
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "tyun.cn:8088")
	if err != nil {
		//handle
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		//读取输入
		readString, _ := inputReader.ReadString('\n')
		trim := strings.Trim(readString, "\r\n")
		_, err := conn.Write([]byte(trim))
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			return
		}
		fmt.Println(string(buf[:n]))
	}

}
