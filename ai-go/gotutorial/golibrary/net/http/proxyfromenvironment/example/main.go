/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 14:10:05
*/
package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

func send_http_request(addr string, port int) error {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	// construct encoded endpoint
	Url, err := url.Parse(fmt.Sprintf("http://%s:%d", addr, port))
	if err != nil {
		return err
	}
	Url.Path += "/index"
	endpoint := Url.String()
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}
	// use httpClient to send request
	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	// close the connection to reuse it
	defer rsp.Body.Close()
	// check status code
	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("get rsp error: %v", rsp)
	}
	return err
}

func main() {
	send_http_request("xxx", 8080)
}
