/*
@File   : main.go
@Author : pan
@Time   : 2025-03-10 16:44:06
*/
package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

type connTracker struct {
	conn net.Conn
}

func (c *connTracker) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	conn, err := (&net.Dialer{}).DialContext(ctx, network, addr)
	if err != nil {
		return nil, err
	}
	c.conn = conn
	return conn, nil
}

func main() {
	tracker := &connTracker{}
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: tracker.DialContext,
		},
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // 禁用重定向
		},
	}

	req, _ := http.NewRequest("GET", "https://www.kancloud.cn", nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")

	// 发起请求并处理结果
	if _, err := client.Do(req); err != nil {
		fmt.Println("解析不了")
		return
	}

	if tracker.conn != nil {
		fmt.Println(tracker.conn.RemoteAddr().String())
	} else {
		fmt.Println("解析不了")
	}
}
