/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:24:43
*/
package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type getBasicAuthTest struct {
	username, password string
	ok                 bool
}

var parseBasicAuthTests = []struct {
	header, username, password string
	ok                         bool
}{
	{"Basic " + base64.StdEncoding.EncodeToString([]byte("Aladdin:open sesame")), "Aladdin", "open sesame", true},

	// 大小写不影响
	{"BASIC " + base64.StdEncoding.EncodeToString([]byte("Aladdin:open sesame")), "Aladdin", "open sesame", true},
	{"basic " + base64.StdEncoding.EncodeToString([]byte("Aladdin:open sesame")), "Aladdin", "open sesame", true},

	{"Basic " + base64.StdEncoding.EncodeToString([]byte("Aladdin:open:sesame")), "Aladdin", "open:sesame", true},
	{"Basic " + base64.StdEncoding.EncodeToString([]byte(":")), "", "", true},
	{"Basic" + base64.StdEncoding.EncodeToString([]byte("Aladdin:open sesame")), "", "", false},
	{base64.StdEncoding.EncodeToString([]byte("Aladdin:open sesame")), "", "", false},
	{"Basic ", "", "", false},
	{"Basic Aladdin:open sesame", "", "", false},
	{`Digest username="Aladdin"`, "", "", false},
}

func main() {
	for _, tt := range parseBasicAuthTests {
		r, _ := http.NewRequest("GET", "http://example.com/", nil)
		r.Header.Set("Authorization", tt.header)
		fmt.Println(r.Header.Get("Authorization")) //得到的是加密后的结果

		username, password, ok := r.BasicAuth()
		if ok != tt.ok || username != tt.username || password != tt.password {
			fmt.Printf("BasicAuth() = %#v, want %#v", getBasicAuthTest{username, password, ok},
				getBasicAuthTest{tt.username, tt.password, tt.ok})
		}
	}
}
