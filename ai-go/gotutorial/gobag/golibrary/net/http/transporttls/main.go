/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 13:03:55
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

func GetHttps(url, caCertPath, certFile, keyFile string) ([]byte, error) {

	// 创建证书池及各类对象
	var pool *x509.CertPool // 我们要把一部分证书存到这个池中
	var client *http.Client
	var resp *http.Response
	var body []byte
	var err error

	var caCrt []byte // 根证书
	caCrt, err = os.ReadFile(caCertPath)
	pool = x509.NewCertPool()
	if err != nil {
		return nil, err
	}
	pool.AppendCertsFromPEM(caCrt)

	var cliCrt tls.Certificate // 具体的证书加载对象
	cliCrt, err = tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// 把上面的准备内容传入 client
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      pool,
				Certificates: []tls.Certificate{cliCrt},
			},
		},
	}

	// Get 请求
	resp, err = client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer client.CloseIdleConnections()
	return body, nil
}

func main() {
	resp, err := GetHttps("https://10.10.102.91:2379/metrics",
		"C:/Users/Liing/Desktop/ca.crt",
		"C:/Users/Liing/Desktop/healthcheck-client.crt",
		"C:/Users/Liing/Desktop/healthcheck-client.key")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(resp))
}
