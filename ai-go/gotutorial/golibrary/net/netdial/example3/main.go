package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	ip := getHostIp()
	fmt.Println(ip)
}

func getHostIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println("get current host ip err: ", err)
		return ""
	}
	addr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(addr.String(), ":")[0]
	return ip
}
