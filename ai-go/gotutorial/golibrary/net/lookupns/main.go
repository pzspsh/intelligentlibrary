/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:10:36
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	nameserver, _ := net.LookupNS("facebook.com")
	for _, ns := range nameserver {
		fmt.Println(ns)
	}
}
