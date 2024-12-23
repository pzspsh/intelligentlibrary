/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 16:40:15
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	response, err := client.Get("http://example.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
