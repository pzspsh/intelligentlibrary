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
	e, err := mail.ParseAddress("Alice <alice@example.com>")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(e.Name, e.Address)
}
