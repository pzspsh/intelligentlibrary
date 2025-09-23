package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodHead, "http://httpbin.org/get", nil)
	if err != nil {
		fmt.Println(err)
	}
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
