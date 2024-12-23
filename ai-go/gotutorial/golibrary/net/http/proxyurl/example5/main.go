/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:39:28
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	proxyURL, _ := url.Parse("http://proxy.example.com:8080")

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{Transport: transport}

	resp, err := client.Get("http://example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))
}
