/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 12:36:59
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

/*
CA证书是不是计算机内置的。那如果我们希望完整的验证过程，那么我们可以把CA证书也发给客户端，客户端代码修改为
*/
func main() {
	pool := x509.NewCertPool()
	caCertPath := "ca.crt"

	caCrt, err := os.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8081")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
