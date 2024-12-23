/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 14:44:20
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	urli := url.URL{}
	urlproxy, _ := urli.Parse("https://127.0.0.1:9743")
	c := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		},
	}
	if resp, err := c.Get("https://www.google.com"); err != nil {
		log.Fatalln(err)
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
	}
}
