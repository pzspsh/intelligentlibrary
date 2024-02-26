/*
@File   : main.go
@Author : pan
@Time   : 2024-02-26 16:45:36
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
	resp, err := client.Get("https://www.example.com")
	if err != nil {
		fmt.Println("client get error:", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io ReadAll error: ", err)
	}
	fmt.Println(string(body))
}
