/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:29:22
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
)

func main() {
	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "this call was relayed by the reverse proxy")
	}))
	defer backendServer.Close()

	rpURL, err := url.Parse(backendServer.URL)
	if err != nil {
		log.Fatal(err)
	}
	frontendProxy := httptest.NewServer(&httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetXForwarded()
			r.SetURL(rpURL)
		},
	})
	defer frontendProxy.Close()

	resp, err := http.Get(frontendProxy.URL)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}
