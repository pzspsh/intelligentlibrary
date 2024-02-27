/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 17:12:09
*/
package main

import (
	"fmt"
	"log"
	"net/smtp"
)

// 邮箱服务器配置信息
type configInof struct {
	smtpAddr string
	smtpPort string
	secret   string
}

// 邮件内容信息
type emailContent struct {
	fromAddr    string
	contentType string
	theme       string
	message     string
	toAddr      []string
}

func sendEmail(c *configInof, e *emailContent) error {
	// 拼接smtp服务器地址
	smtpAddr := c.smtpAddr + ":" + c.smtpPort
	// 认证信息
	auth := smtp.PlainAuth("", e.fromAddr, c.secret, c.smtpAddr)
	// 配置邮件内容类型
	if e.contentType == "html" {
		e.contentType = "Content-Type: text/html; charset=UTF-8"
	} else {
		e.contentType = "Content-Type: text/plain; charset=UTF-8"
	}
	// 当有多个收件人
	for _, to := range e.toAddr {
		msg := []byte("To: " + to + "\r\n" +
			"From: " + e.fromAddr + "\r\n" +
			"Subject: " + e.theme + "\r\n" +
			e.contentType + "\r\n\r\n" +
			"<html><h1>" + e.message + "</h1></html>")
		err := smtp.SendMail(smtpAddr, auth, e.fromAddr, []string{to}, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// 收集配置信息
	config := configInof{
		// smtp服务器地址
		smtpAddr: "smtp.yeah.net",
		// smtp服务器密钥
		secret: "xxxxxxxxxxxxxx",
		// smtp服务器端口
		smtpPort: "25",
	}
	// 收集邮件内容
	content := emailContent{
		// 发件人
		fromAddr: "youremail@yeah.net",
		// 收件人(可有多个)
		toAddr: []string{"xxxxxx@qq.com", "xxxxxxx@126.com"},
		// 邮件格式
		contentType: "html",
		// 邮件主题
		theme: "我是一个正经邮件",
		// 邮件内容
		message: "我有高压锅你要吗",
	}
	// 发送邮件
	err := sendEmail(&config, &content)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("发送成功")
	}
}
