/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:16:36
*/
package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var writeSetCookiesTests = []struct {
	Cookie *http.Cookie
	Raw    string
}{
	{
		&http.Cookie{Name: "cookie-2", Value: "two", MaxAge: 3600},
		"cookie-2=two; Max-Age=3600",
	},
	{
		&http.Cookie{Name: "cookie-3", Value: "three", Domain: ".example.com"},
		"cookie-3=three; Domain=example.com",
	},
	{
		&http.Cookie{Name: "cookie-4", Value: "four", Path: "/restricted/"},
		"cookie-4=four; Path=/restricted/",
	},
	{
		&http.Cookie{Name: "cookie-9", Value: "expiring", Expires: time.Unix(1257894000, 0)},
		"cookie-9=expiring; Expires=Tue, 10 Nov 2009 23:00:00 GMT",
	},
	// According to IETF 6265 Section 5.1.1.5, the year cannot be less than 1601
	{ //故意将这里的cookie-10写成cookie-101，然后下面就会报错
		&http.Cookie{Name: "cookie-10", Value: "expiring-1601", Expires: time.Date(1601, 1, 1, 1, 1, 1, 1, time.UTC)},
		"cookie-101=expiring-1601; Expires=Mon, 01 Jan 1601 01:01:01 GMT",
	},
	{ //因此其返回值中没有Expires
		&http.Cookie{Name: "cookie-11", Value: "invalid-expiry", Expires: time.Date(1600, 1, 1, 1, 1, 1, 1, time.UTC)},
		"cookie-11=invalid-expiry",
	},
	// The "special" cookies have values containing commas or spaces which
	// are disallowed by RFC 6265 but are common in the wild.
	{
		&http.Cookie{Name: "special-1", Value: "a z"},
		`special-1="a z"`,
	},
	{
		&http.Cookie{Name: "empty-value", Value: ""},
		`empty-value=`,
	},
	{
		nil,
		``,
	},
	{
		&http.Cookie{Name: ""},
		``,
	},
	{
		&http.Cookie{Name: "\t"},
		``,
	},
}

/*
Cookie代表一个出现在HTTP回复的头域中Set-Cookie头的值里或者HTTP请求的头域中Cookie头的值里的HTTP cookie。

String返回该cookie的序列化结果。如果只设置了Name和Value字段，序列化结果可用于HTTP请求的Cookie头或者HTTP
回复的Set-Cookie头；如果设置了其他字段，序列化结果只能用于HTTP回复的Set-Cookie头。
*/
func main() {
	defer log.SetOutput(os.Stderr)
	var logbuf bytes.Buffer
	log.SetOutput(&logbuf)

	for i, tt := range writeSetCookiesTests { //没有报错则说明得到的Cookie的值与Raw字符串相等
		if g, e := tt.Cookie.String(), tt.Raw; g != e {
			fmt.Printf("Test %d, expecting:\n%s\nGot:\n%s\n", i, e, g)
			continue
		}
	}
}
