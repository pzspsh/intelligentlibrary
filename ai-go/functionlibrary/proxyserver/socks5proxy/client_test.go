/*
@File   : client_test.go
@Author : pan
@Time   : 2023-11-09 14:20:42
*/
package socket5proxy

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"golang.org/x/net/proxy"
)

func TestProxyClient(t *testing.T) {
	dialer, err := proxy.SOCKS5("tcp", "proxyip:proxyport", nil, proxy.Direct)
	//dc := dialer.(interface {
	//	DialContext(ctx context.Context, network, addr string) (net.Conn, error)
	//})
	//fmt.Println("dialer", dc)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connet to the proxy:", err)
		os.Exit(1)
	}
	tr := &http.Transport{Dial: dialer.Dial}
	httpClient := &http.Client{Transport: tr}
	if resp, err := httpClient.Get("https://www.baidu.com"); err != nil {
	} else {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
	}
}
