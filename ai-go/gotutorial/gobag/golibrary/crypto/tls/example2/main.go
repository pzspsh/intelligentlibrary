/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 12:27:49
*/
package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

// 未完成
func main() {
	// 创建自签名证书
	certPEM := []byte("")
	keyPEM := []byte("")
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		log.Fatal(err)
	}

	// 创建TLS配置
	config := tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// 实现HTTPS服务器
	handler := http.NewServeMux()
	server := http.Server{
		Addr:      ":443",
		Handler:   handler,
		TLSConfig: &config,
	}
	log.Fatal(server.ListenAndServeTLS("", ""))

	// 实现客户端
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost/")
	fmt.Println(resp)
}

func TelnetServer() {
	// Telnet服务器
	config := tls.Config{}
	listener, err := tls.Listen("tcp", ":23", &config)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(conn)
		// ...
	}
}

func TelnetClient() {
	// Telnet客户端
	conn, err := tls.Dial("tcp", "localhost:23", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn)
}
