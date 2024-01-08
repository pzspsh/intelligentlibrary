/*
@File   : main.go
@Author : pan
@Time   : 2024-01-08 12:18:18
*/
package main

import (
	"errors"
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("golangbot123.com")
	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) {
			if dnsErr.Timeout() {
				fmt.Println("operation timed out")
				return
			}
			if dnsErr.Temporary() {
				fmt.Println("temporary error")
				return
			}
			fmt.Println("Generic DNS error", err)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(addr)
}
