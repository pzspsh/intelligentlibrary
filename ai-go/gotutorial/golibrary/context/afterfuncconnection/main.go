/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 16:37:30
*/
package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

/*
这个例子使用了AfterFunc来定义一个从网络中读取的函数。Conn，当上下文被取消时停止读取。
*/
func main() {
	readFromConn := func(ctx context.Context, conn net.Conn, b []byte) (n int, err error) {
		stopc := make(chan struct{})
		stop := context.AfterFunc(ctx, func() {
			conn.SetReadDeadline(time.Now())
			close(stopc)
		})
		n, err = conn.Read(b)
		if !stop() {
			// The AfterFunc was started.
			// Wait for it to complete, and reset the Conn's deadline.
			<-stopc
			conn.SetReadDeadline(time.Time{})
			return n, ctx.Err()
		}
		return n, err
	}

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	conn, err := net.Dial(listener.Addr().Network(), listener.Addr().String())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	b := make([]byte, 1024)
	_, err = readFromConn(ctx, conn, b)
	fmt.Println(err)

}
