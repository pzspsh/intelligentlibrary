/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:19:01
*/
package main

import (
	"fmt"
	"net/http"
)

var newRequestHostTests = []struct {
	in, out string
}{
	{"http://www.example.com/", "www.example.com"},
	{"http://www.example.com:8080/", "www.example.com:8080"},

	{"http://192.168.0.1/", "192.168.0.1"},
	{"http://192.168.0.1:8080/", "192.168.0.1:8080"},
	{"http://192.168.0.1:/", "192.168.0.1"},

	{"http://[fe80::1]/", "[fe80::1]"},
	{"http://[fe80::1]:8080/", "[fe80::1]:8080"},
	{"http://[fe80::1%25en0]/", "[fe80::1%en0]"},
	{"http://[fe80::1%25en0]:8080/", "[fe80::1%en0]:8080"},
	{"http://[fe80::1%25en0]:/", "[fe80::1%en0]"},
}

func main() {
	for i, tt := range newRequestHostTests {
		req, err := http.NewRequest("GET", tt.in, nil)
		if err != nil {
			fmt.Printf("#%v: %v", i, err)
			continue
		}
		if req.Host != tt.out { //返回结果中没有报错，则说明req.Host == tt.out
			fmt.Printf("got %q; want %q", req.Host, tt.out)
		}
	}
}
