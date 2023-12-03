/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 17:03:06
*/
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	pid := os.Getpid()
	log.Printf("%d starting", pid)
	a, err := net.ResolveIPAddr("ip4", "127.0.0.1")
	if err != nil {
		log.Panic(err)
	}
	addr := a
	conn, err := net.ListenIP("ip:253", addr)
	if err != nil {
		log.Panic(err)
	}
	buf := make([]byte, 9999)
	i := 0
	lim := 3
	for {
		n, ra, err := conn.ReadFromIP(buf)
		if err != nil {
			log.Panic(err)
		}
		s := string(buf[:n])
		s = strings.TrimSpace(s)
		log.Printf("received from %s: %s", ra.String(), s)
		if i < lim {
			i++
			go send(fmt.Sprintf("%d:%d(%s)", pid, i, s))
		}
	}
}

func send(sen string) {

}
