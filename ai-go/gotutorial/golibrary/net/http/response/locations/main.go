/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:36:26
*/
package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type responseLocationTest struct {
	location string // Response's Location header or ""
	requrl   string // Response.Request.URL or ""
	want     string
	wantErr  error
}

var responseLocationTests = []responseLocationTest{
	{"/foo", "http://bar.com/baz", "http://bar.com/foo", nil},
	{"http://foo.com/", "http://bar.com/baz", "http://foo.com/", nil},
	{"", "http://bar.com/baz", "", http.ErrNoLocation},
	{"/bar", "", "/bar", nil},
}

func main() {
	for i, tt := range responseLocationTests {
		res := new(http.Response)
		res.Header = make(http.Header)
		res.Header.Set("Location", tt.location)
		if tt.requrl != "" {
			res.Request = &http.Request{}
			var err error
			res.Request.URL, err = url.Parse(tt.requrl)
			if err != nil {
				fmt.Printf("bad test URL %q: %v", tt.requrl, err)
			} else {
				fmt.Println(i, "URL : ", res.Request.URL)
			}

		}

		got, err := res.Location()
		if tt.wantErr != nil {
			if err == nil {
				fmt.Printf("%d. err=nil; want %q", i, tt.wantErr)
				continue
			}
			if g, e := err.Error(), tt.wantErr.Error(); g != e {
				fmt.Printf("%d. err=%q; want %q", i, g, e)
				continue
			} else {
				fmt.Println(i, "err : ", err.Error())
			}
			continue
		}
		if err != nil {
			fmt.Printf("%d. err=%q", i, err)
			continue
		}
		if g, e := got.String(), tt.want; g != e {
			fmt.Printf("%d. Location=%q; want %q", i, g, e)
		} else {
			fmt.Println(i, "got : ", got.String())
		}
	}
}
