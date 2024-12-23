/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:21:26
*/
package main

import (
	"fmt"
	"net/http"
)

var addCookieTests = []struct {
	Cookies []*http.Cookie
	Raw     string
}{
	{
		[]*http.Cookie{},
		"",
	},
	{
		[]*http.Cookie{{Name: "cookie-1", Value: "v$11"}},
		"cookie-1=v$11",
	},
	{
		[]*http.Cookie{
			{Name: "cookie-1", Value: "v$21"},
			{Name: "cookie-2", Value: "v$2"},
			{Name: "cookie-3", Value: "v$3"},
		},
		"cookie-1=v$21; cookie-2=v$2; cookie-3=v$3",
	},
}

func main() {
	for i, tt := range addCookieTests {
		req, _ := http.NewRequest("GET", "http://example.com/", nil)
		for _, c := range tt.Cookies {
			req.AddCookie(c)
		}
		//没有报错，则说明添加进的Cookie的值与给的Raw的字符串的值相同
		//得到Cookie的值可以使用req.Header.Get("Cookie")，也可以使用下面的req.Cookies()
		if g := req.Header.Get("Cookie"); g != tt.Raw {
			fmt.Printf("Test %d:\nwant: %s\n got: %s\n", i, tt.Raw, g)
			continue
		}
		fmt.Println(req.Cookies())
		value, _ := req.Cookie("cookie-1")
		fmt.Println(value)

	}
}
