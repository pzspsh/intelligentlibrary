/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:39:19
*/
package main

import (
	"fmt"
	"log"
	"net/mail"
)

func main() {
	const list = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
	emails, err := mail.ParseAddressList(list)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range emails {
		fmt.Println(v.Name, v.Address)
	}
}
