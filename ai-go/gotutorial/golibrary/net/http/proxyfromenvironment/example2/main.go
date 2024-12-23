/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 14:31:39
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

/*
export HTTP_PROXY=http://proxy.example.com:8080
export NO_PROXY=localhost,127.0.0.1

os.Setenv("HTTP_PROXY", "http://proxy.example.com:8080")
os.Setenv("NO_PROXY", "localhost,127.0.0.1")
*/
