/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 17:20:16
*/
package main

import (
	"log"
	"net/smtp"
)

const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = ":587"
	SMTPUsername = "xxx@gmail.com"
	SMTPPassword = "xxxx"
)

func SendEmail(receiver string) {
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	msg := []byte("Subject: 这里是标题内容\r\n\r\n" + "这里是正文内容\r\n")
	err := smtp.SendMail(SMTPHost+SMTPPort, auth, SMTPUsername, []string{receiver}, msg)
	if err != nil {
		log.Fatal("failed to send email:", err)
	}
}

func SendHTMLEmail(receiver string, html []byte) {
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	msg := append([]byte("Subject: 这里是标题内容\r\n"+
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"),
		html...)
	err := smtp.SendMail(SMTPHost+SMTPPort, auth, SMTPUsername, []string{receiver}, msg)
	if err != nil {
		log.Fatal("failed to send email:", err)
	}
}

func main() {
	SendHTMLEmail("接受者@gmail.com", []byte("<html><h2>这是网页内容</h2></html>"))
}
