/*
@File   : mian.go
@Author : pan
@Time   : 2024-08-01 11:46:58
*/
package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func init() {
	os.Setenv("GODEBUG", "x509ignoreCN=0")
}

func main() {

	//代理地址
	parse, err := url.Parse("https://www.open1.com:2001")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(parse)

	// 这里是配置了/etc/hosts www.open1.com 127.0.0.1
	//由于要代理https(还是希望通过https访问 （https://127.0.0.1:2000 ｜ https://www.open1.com:2000)代理到 https://www.open1.com:2001)
	//就不能使用默认的http.DefaultTransport

	//如果不希望通过https代理，可以使用默认的 http.DefaultTransport 或者设置 TLSClientConfig: &tls.Config{InsecureSkipVerify: true}
	//这样就能通过 （http://127.0.0.1:2000 ｜ http://www.open1.com:2000)访问了
	proxy.Transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   20 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		//跳过认证
		//TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		TLSClientConfig: func() *tls.Config {
			pool := x509.NewCertPool()
			file, _ := os.ReadFile("/path/ca.crt")
			pool.AppendCertsFromPEM(file)
			return &tls.Config{RootCAs: pool}
		}(),
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", proxy.ServeHTTP)
	server := &http.Server{
		Addr:         ":2000",
		Handler:      mux,
		WriteTimeout: time.Second * 3,
	}

	//传入ssl证书和服务器私钥
	//非https的使用 server.ListenAndServe() 启动
	log.Fatal(server.ListenAndServeTLS("/path/server.crt", "/path/server.key"))
}

// GODEBUG="x509ignoreCN=0" go run main.go
