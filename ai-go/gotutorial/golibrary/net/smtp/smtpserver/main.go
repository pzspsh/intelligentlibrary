/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 17:15:15
*/
package main

import (
	"fmt"
	"net"
	"os"
)

type Mail struct {
	clientName string
	from       string
	to         string
	data       string
	reciveData bool
}

func main() {

	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP("127.0.0.1"), 25, ""})
	if err != nil {
		fmt.Println("监听端口失败:", err.Error())
		return
	}
	defer listen.Close()
	recive(listen)

}

func recive(listen *net.TCPListener) {

	for {
		conn, err := listen.AcceptTCP() //接收连接
		defer conn.Close()
		if err != nil {
			fmt.Println("客户端连接异常", err.Error()) //客户端连接异常
			continue
		}
		conn.Write([]byte("220 GoMailSever "))
		fmt.Println("新客户端", conn.RemoteAddr().String())
		go func() {

			tmpBuffer := make([]byte, 0)           //创建待处理字节切片
			readerChannel := make(chan []byte, 16) //创建channel通信管道
			_mail := &Mail{}
			go reader(readerChannel, conn, _mail) //channel信息接收端
			data := make([]byte, 1024)            //创建输出缓冲切片

			for {
				i, err := conn.Read(data) //从Socket读取数据
				if err != nil {
					fmt.Println("读取发生错误", err.Error()) //读取失败发生错误
					return
				}
				tmpBuffer = DepackMail(append(tmpBuffer, data[:i]...), readerChannel) //用自定义协议处理字节

				//fmt.Println(data, string(data))
			}

		}()

	}
}
func reader(readerChannel chan []byte, conn *net.TCPConn, _mail *Mail) {
	for {
		select {
		case data := <-readerChannel:
			fmt.Println(string(data[0:4]))
			switch string(data[0:4]) {
			case "HELO":
				_mail.clientName = string(data[4+2:])
				conn.Write([]byte("250 Give Me Mail"))
				break
			case "MAIL":
				_mail.from = string(data[4+2:])
				conn.Write([]byte("250 Give I am fine"))
				break
			case "RCPT":
				_mail.to = string(data[4+2:])
				conn.Write([]byte("250 Give I am fine"))
				break
			case "DATA":
				conn.Write([]byte("354 Give I am fine"))
				break
			case "QUIT":
				conn.Write([]byte("221 ByeBye"))
				file, _ := os.Create("./mail/3.eml")
				file.WriteString(_mail.data)
				file.Close()
				fmt.Println(_mail)
				break
			default:
				_mail.data = _mail.data + string(data)
				fmt.Println(string(_mail.data[len(_mail.data)-1:]))
				if string(_mail.data[len(_mail.data)-1:]) == "." {
					fmt.Println("111")
					conn.Write([]byte("250 Give I am fine"))
				}
				break
			}
		}
	}
}
func DepackMail(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)

	var i int

	for i = 0; i < length; {
		if length < i+4 { //2位为尾标记
			break
		}
		var sIndex int
		for dump := i; dump < length-1; dump++ {
			if buffer[dump] == byte(13) && buffer[dump+1] == byte(10) {

				sIndex = dump
			}
		}
		if sIndex <= 0 {
			break
		}
		data := buffer[i:sIndex]

		readerChannel <- data //向管道发送数据
		i = sIndex + 2
	}

	if i == length { //全部处理完毕返回空包
		return make([]byte, 0)
	}
	return buffer[i:] //未处理完毕返回 剩余包
}
