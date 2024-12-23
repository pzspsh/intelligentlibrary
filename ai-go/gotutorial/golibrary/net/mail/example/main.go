/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:57:18
*/
package main

import (
	"fmt"
	"net/mail"
	"time"
)

func main() {
	dateString := "Fri, 24 Mar 2023 10:00:00 GMT"
	parsedDate, err := mail.ParseDate(dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	fmt.Printf("Parsed date: %s\n", parsedDate.Format(time.RFC1123))
}
