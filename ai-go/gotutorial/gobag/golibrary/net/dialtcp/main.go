/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 19:36:22
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	// Set the timeout to 5 seconds
	// timeout := 5 * time.Second

	// Create a Dialer with the timeout set
	// dialer := net.Dialer{
	// 	Timeout: timeout,
	// }

	// DialTCP with the Dialer
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Do something with the connection
	fmt.Println("Connection established:", conn)
}
