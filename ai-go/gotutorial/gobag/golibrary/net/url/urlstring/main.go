/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 17:30:16
*/
package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := &url.URL{
		Scheme:   "https",
		User:     url.UserPassword("me", "pass"),
		Host:     "example.com",
		Path:     "foo/bar",
		RawQuery: "x=1&y=2",
		Fragment: "anchor",
	}
	fmt.Println(u.String())
	u.Opaque = "opaque"
	fmt.Println(u.String())
}
