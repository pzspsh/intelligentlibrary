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
