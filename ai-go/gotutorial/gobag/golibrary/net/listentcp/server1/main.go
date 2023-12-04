/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 10:59:05
*/
package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":4044")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen to 4044")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("conn err:", err)
		} else {
			go handleConn2(conn)
		}
	}
}

func packetSlitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 检查 atEOF 参数 和 数据包头部的四个字节是否 为 0x123456(我们定义的协议的魔数)
	if !atEOF && len(data) > 6 && binary.BigEndian.Uint32(data[:4]) == 0x123456 {
		var l int16
		// 读出 数据包中 实际数据 的长度(大小为 0 ~ 2^16)
		binary.Read(bytes.NewReader(data[4:6]), binary.BigEndian, &l)
		pl := int(l) + 6
		if pl <= len(data) {
			return pl, data[:pl], nil
		}
	}
	return
}

func handleConn2(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("关闭")
	fmt.Println("新连接：", conn.RemoteAddr())
	result := bytes.NewBuffer(nil)
	var buf [65542]byte // 由于 标识数据包长度 的只有两个字节 故数据包最大为 2^16+4(魔数)+2(长度标识)
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				fmt.Println("read err:", err)
				break
			}
		} else {
			scanner := bufio.NewScanner(result)
			scanner.Split(packetSlitFunc)
			for scanner.Scan() {
				fmt.Println("recv:", string(scanner.Bytes()[6:]))
			}
		}
		result.Reset()
	}
}
