package main

import (
	"fmt"
	"function/ftp/ftp"
	"time"
)

type FTP struct {
	Host     string
	Port     int64
	Proxy    string
	Username string // 用户名
	Password string // 密码
	Cmd      string
}

func (f *FTP) FtpConnect() {
	addr := fmt.Sprintf("%v:%v", f.Host, f.Port)
	servConn, err := ftp.Dial(addr, f.Proxy, ftp.DialWithTimeout(time.Duration(time.Second*2)*time.Second))
	if err != nil {
		return
	}
	err = servConn.Login(f.Username, f.Password)
	if err != nil {
		fmt.Println("连接失败")
	} else {
		fmt.Println("连接成功")
	}
}

func main() {

}
