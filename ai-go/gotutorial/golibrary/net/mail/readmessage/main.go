/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:39:19
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/mail"
	"strings"
)

func main() {
	msg := `Date: Mon, 23 Jun 2015 11:40:36 -0400
From: Gopher <from@example.com>
To: Another Gopher <to@example.com>
Subject: Gophers at Gophercon

Message body
`

	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	header := m.Header
	fmt.Println("Date:", header.Get("Date"))
	fmt.Println("From:", header.Get("From"))
	fmt.Println("To:", header.Get("To"))
	fmt.Println("Subject:", header.Get("Subject"))

	body, err := io.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

}
