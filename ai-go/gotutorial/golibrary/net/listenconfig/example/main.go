/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:00:48
*/
package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {
	lc := net.ListenConfig{}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("cancelling context...")
		cancel()
	}()
	ln, err := lc.Listen(ctx, "tcp", ":9801")
	if err != nil {
		fmt.Println("error creating listener:", err)
	} else {
		fmt.Println("listen returned without error")
		defer ln.Close()
	}
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("accept returned error:", err)
	} else {
		fmt.Println("accept returned without error")
		defer conn.Close()
	}
}
