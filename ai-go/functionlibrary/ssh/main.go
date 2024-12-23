package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
	"golang.org/x/net/proxy"
)

type SSH struct {
	Host     string
	Port     int64
	Proxy    string
	Username string // 用户名
	Password string // 密码
	Cmd      string
}

func (s *SSH) Sshconnect() {
	sshConfig := &ssh.ClientConfig{
		User:            s.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(s.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	var ret bool
	addr := fmt.Sprintf("%s:%v", s.Host, s.Port)
	if s.Proxy != "" {
		dialer, err := proxy.SOCKS5("tcp", s.Proxy, nil, proxy.Direct)
		if err != nil {
			ret = false
		}
		conn, err := dialer.Dial("tcp", addr)
		if err != nil {
			ret = false
		}
		_, _, _, err = ssh.NewClientConn(conn, addr, sshConfig)
		if err != nil {
			ret = false
		} else {
			ret = true
		}
	} else {
		_, err := ssh.Dial("tcp", addr, sshConfig)
		if err != nil {
			ret = false
		} else {
			ret = true
		}
	}
	if ret {
		fmt.Println("连接成功")
	}
}

func main() {

}
