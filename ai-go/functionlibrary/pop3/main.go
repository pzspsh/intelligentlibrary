package main

import (
	"fmt"
	"function/pop3/pop3"
)

type POP3 struct {
	Host     string
	Port     int64
	Proxy    string
	Username string // 用户名
	Password string // 密码
	Cmd      string
}

func (p *POP3) Pop3Connet() {
	po := pop3.Opt{}
	if p.Proxy != "" {
		po.Proxy = p.Proxy
	}
	po.Host = p.Host
	po.Port = int(p.Port)
	po.TLSEnabled = true
	pop := pop3.New(po)
	c, err := pop.NewConn()
	if err != nil {
		return
	}
	defer c.Quit()

	// Authenticate.
	err = c.Auth(p.Username, p.Password)
	if err != nil {
		fmt.Println("连接失败")
	} else {
		fmt.Println("连接成功")
	}
}

func main() {

}
