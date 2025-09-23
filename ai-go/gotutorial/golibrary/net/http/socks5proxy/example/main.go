package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/proxy"
)

func main() {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:8080",
		&proxy.Auth{User: "username", Password: "password"},
		&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	trans := &http.Transport{
		Dial: dialer.Dial,
	}
	client := &http.Client{
		Transport: trans,
	}
	client.Get("https://www.google.com")
}
