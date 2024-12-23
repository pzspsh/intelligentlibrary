/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 17:09:48
*/
package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial("mail.example.com:25")
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail("sender@example.org"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("recipient@example.net"); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
