/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 19:59:32
*/
package main

import "net"

func main() {
	src, err := net.Dial("udp4", "238.98.76.54:4378")
	if err != nil {
		panic("test failed")
	}
	defer src.Close()
	addr := net.ParseIP("238.98.76.54")
	dest, err := net.ListenMulticastUDP("udp4", nil, &net.UDPAddr{IP: addr, Port: 4378})
	if err != nil {
		panic("test failed")
	}
	defer dest.Close()
	if _, err := src.Write(make([]byte, 100)); err != nil {
		panic("test failed")
	}
	b := make([]byte, 1024)
	n, _, err := dest.ReadFrom(b)
	if err != nil || n != 100 {
		panic("test failed")
	}
}
