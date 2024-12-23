/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:58:27
*/
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// create cert pool, which stands for cert set
	pool := x509.NewCertPool()
	caCrtPath := "certpath/client/ca.crt"
	// call ca.crt
	caCrt, err := os.ReadFile(caCrtPath)
	if err != nil {
		log.Println("ReadFile err:", err)
		return
	}

	// parse cert
	pool.AppendCertsFromPEM(caCrt)
	tr := &http.Transport{
		// InsecureSkipVerify-如果设置为true, 则不会校验证书以及证书中的主机名和服务器主机名是否一致
		TLSClientConfig: &tls.Config{RootCAs: pool, InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://localhost:8443/index")
	if err != nil {
		log.Println("Get err:", err)
		return
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(string(content))
}
