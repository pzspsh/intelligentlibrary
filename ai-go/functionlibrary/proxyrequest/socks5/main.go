package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/proxy"
)

func Socks5Set() {
	// 创建一个代理Dialer
	dialer, err := proxy.SOCKS5("tcp", "proxy_address:port", nil, proxy.Direct) // "proxy_address:port"是ip:port格式
	if err != nil {
		panic(err)
	}

	// 使用代理Dialer创建一个自定义的Transport
	transport := &http.Transport{
		Dial: dialer.Dial, // 使用代理Dialer代替默认的Dialer
	}

	// 使用自定义Transport创建一个Client
	client := &http.Client{
		Transport: transport,
	}

	// 使用Client发起HTTP请求
	resp, err := client.Get("http://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应内容并打印
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func Socks5Set1() {
	// 定义 SOCKS5 代理地址
	socks5Addr := "127.0.0.1:1080" // 替换为实际的代理地址

	// 创建自定义 Dial 函数
	dialer := func(network, addr string) (net.Conn, error) {
		// 连接到 SOCKS5 代理
		conn, err := net.Dial("tcp", socks5Addr)
		if err != nil {
			return nil, err
		}

		// 发送 SOCKS5 握手请求
		_, err = conn.Write([]byte{0x05, 0x01, 0x00}) // 协议版本 5，1 种认证方法，无认证
		if err != nil {
			conn.Close()
			return nil, err
		}

		// 读取握手响应
		buf := make([]byte, 2)
		_, err = conn.Read(buf)
		if err != nil || buf[0] != 0x05 || buf[1] != 0x00 {
			conn.Close()
			return nil, fmt.Errorf("SOCKS5 握手失败")
		}

		// 发送 CONNECT 请求
		dest := fmt.Sprintf("%s\x00", addr)
		req := []byte{0x05, 0x01, 0x00, 0x03, byte(len(dest))}
		req = append(req, []byte(dest)...)

		_, err = conn.Write(req)
		if err != nil {
			conn.Close()
			return nil, err
		}

		// 读取 CONNECT 响应
		buf = make([]byte, 4)
		_, err = conn.Read(buf)
		if err != nil || buf[1] != 0x00 {
			conn.Close()
			return nil, fmt.Errorf("SOCKS5 CONNECT 失败")
		}

		return conn, nil
	}

	// 创建自定义 HTTP 客户端
	httpTransport := &http.Transport{
		Dial: dialer,
	}
	client := &http.Client{
		Transport: httpTransport,
	}

	// 发送请求
	resp, err := client.Get("http://example.com")
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应失败: %v", err)
	}

	fmt.Println(string(body))
}

func Socks5Set2() {
	// 设置环境变量
	// export HTTP_PROXY=socks5://127.0.0.1:1080
	// export HTTPS_PROXY=socks5://127.0.0.1:1080

	// 创建默认的 HTTP 客户端
	client := &http.Client{}

	// 发送请求
	resp, err := client.Get("http://example.com")
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应失败: %v", err)
	}

	fmt.Println(string(body))
}

func Socks5Set3() {
	var HTTP_PROXY_ENV = "HTTP_PROXY"
	aliveProxy := "socks5://127.0.0.1:1080"
	proxyURL, err := url.Parse(aliveProxy)
	if err != nil {
		log.Fatalf("解析代理地址失败: %v", err)
	}
	os.Setenv(HTTP_PROXY_ENV, proxyURL.String())
}

func Socks5Set4() {
	// 定义 SOCKS5 代理地址
	socks5Proxy := "socks5://127.0.0.1:1080" // 替换为实际的代理地址
	proxyURL, err := url.Parse(socks5Proxy)
	if err != nil {
		log.Fatalf("解析代理地址失败: %v", err)
	}
	// 解析代理 URL
	proxyurl, err := proxy.FromURL(proxyURL, proxy.Direct)
	if err != nil {
		log.Fatalf("解析代理失败: %v", err)
	}

	// 创建自定义的 HTTP 客户端
	httpTransport := &http.Transport{
		Dial: proxyurl.Dial,
	}
	client := &http.Client{
		Transport: httpTransport,
	}

	// 发送请求
	resp, err := client.Get("http://example.com")
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应失败: %v", err)
	}

	fmt.Println(string(body))
}

func Socks5Set5() {
	// 定义 SOCKS5 代理地址
	socks5Proxy := "socks5://127.0.0.1:1080" // 替换为实际的代理地址
	proxyURL, err := url.Parse(socks5Proxy)
	if err != nil {
		log.Fatalf("解析代理地址失败: %v", err)
	}
	// 解析代理 URL
	dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
	if err != nil {
		log.Fatalf("解析代理失败: %v", err)
	}
	tlsConfig := &tls.Config{
		Renegotiation:      tls.RenegotiateOnceAsClient,
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS10,
	}
	transport := &http.Transport{}
	dc := dialer.(interface {
		DialContext(ctx context.Context, network, addr string) (net.Conn, error)
	})
	transport.DialContext = dc.DialContext
	transport.DialTLSContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		// upgrade proxy connection to tls
		conn, err := dc.DialContext(ctx, network, addr)
		if err != nil {
			return nil, err
		}
		return tls.Client(conn, tlsConfig), nil
	}
	client := &http.Client{Transport: transport}
	fmt.Println(client)
}

func main() {

}
