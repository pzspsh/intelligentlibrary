package main

import (
	"fmt"
	"net/url"
)

func main() {
	path := url.PathEscape("my/cool+blog&about,stuff")
	fmt.Println(path)

}
