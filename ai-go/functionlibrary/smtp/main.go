/*
@File   : smtp.go
@Author : pan
@Time   : 2023-06-01 15:55:58
*/
package main

import (
	"fmt"
	"function/smtp/smtp"
)

type SMTP struct {
	Host     string
	Port     int64
	Proxy    string
	Username string // 用户名
	Password string // 密码
	Cmd      string
}

func (s *SMTP) SmtpConnect() {
	// auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	err := smtp.SendMail(fmt.Sprintf("%s%b", s.Host, s.Port), s.Proxy, auth, s.Username, []string{}, []byte(""))
	if err != nil {
		fmt.Println("连接失败")
	} else {
		fmt.Println("连接成功")
	}
}

func main() {

}
