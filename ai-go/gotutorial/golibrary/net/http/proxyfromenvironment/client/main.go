/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 14:14:13
*/
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	tr, _ := TLSTransport("/home/ao/Documents/certs/review/server.crt")
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", "https://test.openssl.com:1213/https", nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("err: %+v", err)
	} else {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("resp: %+v=>%+v", resp.StatusCode, string(body))
	}
}

func TLSTransport(caFile string) (*http.Transport, error) {
	tr := &http.Transport{TLSClientConfig: &tls.Config{}, Proxy: http.ProxyFromEnvironment}
	if len(caFile) == 0 {
		tr.TLSClientConfig.InsecureSkipVerify = true
		return tr, nil
	}

	ca, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("read CA file failed: %v", err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(ca)

	tr.TLSClientConfig.RootCAs = pool

	return tr, nil
}
