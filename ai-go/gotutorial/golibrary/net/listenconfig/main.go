/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 17:56:51
*/
package main

// import (
// 	"context"
// 	"fmt"
// 	"net"
// 	"syscall"
// )

// func serve(addr string) {
// 	cfg := net.ListenConfig{
// 		Control: func(network, address string, c syscall.RawConn) error {
// 			return c.Control(func(fd uintptr) {
// 				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEADDR, 1)
// 			})
// 		},
// 	}
// 	udp, err := cfg.ListenPacket(context.Background(), "udp", addr)

// 	if err != nil {
// 		fmt.Println("listen failed", err)
// 		return
// 	}

// 	buf := make([]byte, 1024)
// 	for {
// 		n, caddr, err := udp.ReadFrom(buf)
// 		if err != nil {
// 			fmt.Println("read failed", err)
// 			continue
// 		}

// 		fmt.Println(addr, caddr, string(buf[:n]))
// 	}
// }

// func main() {
// 	go serve("127.0.0.1:1234")
// 	go serve("0.0.0.0:1234")
// 	select {}
// }
