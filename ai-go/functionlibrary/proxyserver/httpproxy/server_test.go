/*
@File   : server_test.go
@Author : pan
@Time   : 2023-11-09 14:41:34
*/
package httpproxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"testing"
)

func TestStartHttp(t *testing.T) {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "https",
		Host:   "www.google.com",
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: proxy,
	}
	log.Fatal(server.ListenAndServe())
}
