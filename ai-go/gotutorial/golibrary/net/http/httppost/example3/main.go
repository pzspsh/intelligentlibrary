package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	targetUrl := "https://blog.csdn.net/zyndev"
	payload := strings.NewReader("a=111")
	response, err := http.Post(targetUrl, "x-www-form-urlencoded", payload)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	fmt.Println(response)
}
