/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:04:38
*/
package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	from := "sender@example.com"
	to := []string{"recipient@example.com"}
	subject := "Subject: Test email from Go\n"
	body := "Hello, this is a test email sent using Go's smtp package."
	msg := []byte(subject + "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n" + body)
	err := smtp.SendMail(
		"smtp.example.com:587",
		smtp.PlainAuth("", "user@example.com", "password", "smtp.example.com"),
		from, to, msg)

	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
