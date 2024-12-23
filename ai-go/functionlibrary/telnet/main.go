package main

import (
	"fmt"
	"function/telnet/telnet"
	"net"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

type Telnet struct {
	Host     string
	Port     int64
	Proxy    string
	Username string // 用户名
	Password string // 密码
	Cmd      string
}

func (t *Telnet) TelnetConnect() {
	var conn net.Conn
	var err error
	addr := fmt.Sprintf("%v:%v", t.Host, t.Port)
	if t.Proxy != "" {
		dialer, _ := proxy.SOCKS5("tcp", t.Proxy, nil, proxy.Direct)
		conn, err = dialer.Dial("tcp", addr)
		if err != nil {
			return
		}
	} else {
		conn, err = net.DialTimeout("tcp", addr, time.Duration(5)*time.Second)
		if err != nil {
			return
		}
	}

	defer conn.Close()
	if _, err := telnet.TelnetProtocolConn(conn); err != nil {
		return
	}
	data := make([]byte, 1024)
	_, err = conn.Read(data)
	if err != nil {
		return
	}
	_, err = conn.Write([]byte(t.Username + "\n"))
	if err != nil {
		return
	}
	_, err = conn.Read(data)
	if err != nil {
		return
	}
	_, err = conn.Read(data)
	if err != nil {
		return
	}
	_, err = conn.Write([]byte(t.Password + "\n"))
	if err != nil {
		return
	}
	_, err = conn.Read(data)
	if err != nil {
		return
	}
	_, err = conn.Read(data)
	if err != nil {
		return
	} else if strings.Contains(string(data), "Last login:") {
		fmt.Println("连接成功")
	}
}

func main() {

}
