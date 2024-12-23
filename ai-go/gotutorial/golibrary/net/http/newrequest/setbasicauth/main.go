/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:23:45
*/
package main

import (
	"fmt"
	"net/http"
)

type getBasicAuthTest struct {
	username, password string
	ok                 bool
}

type basicAuthCredentialsTest struct {
	username, password string
}

var getBasicAuthTests = []struct {
	username, password string
	ok                 bool
}{
	{"Aladdin", "open sesame", true},
	{"Aladdin", "open:sesame", true},
	{"", "", true},
}

func main() {
	for _, tt := range getBasicAuthTests {
		r, _ := http.NewRequest("GET", "http://example.com/", nil)
		r.SetBasicAuth(tt.username, tt.password)
		fmt.Println(r.Header.Get("Authorization")) //在Header中授权信息是加密过的，返回：
		// Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==
		// Basic QWxhZGRpbjpvcGVuOnNlc2FtZQ==
		// Basic Og==

		username, password, ok := r.BasicAuth()
		if ok != tt.ok || username != tt.username || password != tt.password { //满足其中的任意一种情况都说明有错
			fmt.Printf("BasicAuth() = %#v, want %#v", getBasicAuthTest{username, password, ok},
				getBasicAuthTest{tt.username, tt.password, tt.ok})
		}
	}

	//没有授权的request
	r, _ := http.NewRequest("GET", "http://example.com/", nil)
	username, password, ok := r.BasicAuth()
	fmt.Println(username, password, ok) //因为没有授权，返回 "" "" false
	if ok {
		fmt.Printf("expected false from BasicAuth when the request is unauthenticated")
	}
	want := basicAuthCredentialsTest{"", ""} //没有授权返回的username和password都应该为""
	if username != want.username || password != want.password {
		fmt.Printf("expected credentials: %#v when the request is unauthenticated, got %#v",
			want, basicAuthCredentialsTest{username, password})
	}
}
