package main

import (
	"fmt"
	"net/url"
)

func main() {
	query := url.QueryEscape("my/cool+blog&about,stuff")
	fmt.Println(query)

}
