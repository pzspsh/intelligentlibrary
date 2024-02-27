/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:09:47
*/
package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func main() {
	cookies := make([]*http.Cookie, 0)
	cookies = append(cookies, &http.Cookie{
		Name:   "name",
		Value:  "poloxue",
		Domain: "httpbin.org",
		Path:   "/cookies",
	})
	cookies = append(cookies, &http.Cookie{
		Name:   "id",
		Value:  "10000",
		Domain: "httpbin.org",
		Path:   "/elsewhere",
	})

	url, err := url.Parse("http://httpbin.org/cookies")
	if err != nil {
		panic(err)
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	jar.SetCookies(url, cookies)

	client := http.Client{Jar: jar}

	r, err := client.Get("http://httpbin.org/cookies")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
