/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 14:45:18
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/proxy"
)

func main() {
	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9742", nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}

	// set our socks5 as the dialer
	httpTransport := &http.Transport{
		Dial: dialer.Dial,
	}
	// setup a http client
	httpClient := &http.Client{Transport: httpTransport}

	if resp, err := httpClient.Get("https://www.google.com"); err != nil {
		log.Fatalln(err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
	}
}
