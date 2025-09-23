package main

import (
	"fmt"
	"net/url"
)

func main() {
	params := url.Values{}
	params.Add("message", "this will be esc@ped!")
	params.Add("author", "golang c@fe >.<")
	fmt.Println("http://example.com/say?" + params.Encode())
}
