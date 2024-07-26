/*
@File   : main.go
@Author : pan
@Time   : 2024-07-26 15:20:51
*/
package main

import (
	"context"
	"fmt"
	"net"
	"net/url"

	"golang.org/x/net/proxy"
)

// Socks5Dialer 返回一个通过socks5代理连接到目标地址的Dialer
func Socks5Dialer(proxyAddr string, target string) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		// 连接到SOCKS5代理服务器
		var proxyAuth *proxy.Auth
		var Conn net.Conn
		var err error
		socksURL, proxyErr := url.Parse(proxyAddr)
		// 通过SOCKS5代理连接到目标地址
		if proxyErr == nil {
			proxyAuth = &proxy.Auth{}
			proxyAuth.User = socksURL.User.Username()
			proxyAuth.Password, _ = socksURL.User.Password()
		}
		dialer, proxyEr := proxy.SOCKS5("tcp", fmt.Sprintf("%s:%s", socksURL.Hostname(), socksURL.Port()), proxyAuth, proxy.Direct)
		dc := dialer.(interface {
			DialContext(ctx context.Context, network, addr string) (net.Conn, error)
		})
		if proxyEr == nil {
			Conn, err = dc.DialContext(context.Background(), "tcp", "")
			if err != nil {
				return nil, err
			}
			return Conn, nil
		}
		return nil, err
		// 注意：这里我们没有关闭proxyConn，因为它是由targetConn管理的
		// 如果你需要显式控制，可能需要修改库的实现或自己的封装
	}
}

func main() {
	// 设置SOCKS5代理服务器的地址
	proxyAddr := "socks5://219.150.218.53:40154"

	// 使用自定义的Socks5 Dialer
	dialer := Socks5Dialer(proxyAddr, "www.baidu.com:80")

	// 连接到目标服务器（通过SOCKS5代理）
	conn, err := dialer(context.Background(), "tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()

	// 现在你可以使用conn进行通信了
}
