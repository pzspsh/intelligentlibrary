/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 15:39:49
*/
package main

import (
	"context"
	"log"
	"net"
	"time"
)

func main() {
	// DialUnix does not take a context.Context parameter. This example shows
	// how to dial a Unix socket with a Context. Note that the Context only
	// applies to the dial operation; it does not apply to the connection once
	// it has been established.
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	d.LocalAddr = nil // if you have a local addr, add it here
	raddr := net.UnixAddr{Name: "/path/to/unix.sock", Net: "unix"}
	conn, err := d.DialContext(ctx, "unix", raddr.String())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	if _, err := conn.Write([]byte("Hello, socket!")); err != nil {
		log.Fatal(err)
	}
}
