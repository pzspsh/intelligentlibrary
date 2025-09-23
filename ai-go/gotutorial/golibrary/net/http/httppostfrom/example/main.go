package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	targetUrl := "https://blog.csdn.net/zyndev"
	payload := url.Values{"key": {"value"}, "id": {"123"}}
	response, err := http.PostForm(targetUrl, payload)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response)
}
