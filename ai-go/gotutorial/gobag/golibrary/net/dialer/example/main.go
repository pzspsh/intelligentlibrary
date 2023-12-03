/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 15:53:33
*/
package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"syscall"
)

func customControl(network string, address string, conn syscall.RawConn) error {
	fmt.Printf("%s,%s\n", network, address)
	return nil
}

func main() {
	dialer := &net.Dialer{
		Control: customControl,
	}
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.DialContext = dialer.DialContext

	c := &http.Client{
		Transport: t,
	}
	resp, _ := c.Get("https://www.baidu.com")
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)
}
