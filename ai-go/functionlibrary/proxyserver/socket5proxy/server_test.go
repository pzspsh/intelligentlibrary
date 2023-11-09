/*
@File   : server_test.go
@Author : pan
@Time   : 2023-11-09 11:18:52
*/
package socket5proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"golang.org/x/net/proxy"
)

func TestProxyClient(t *testing.T) {
	dialer, err := proxy.SOCKS5("tcp", "10.0.26.11:8081", nil, proxy.Direct)
	//dc := dialer.(interface {
	//	DialContext(ctx context.Context, network, addr string) (net.Conn, error)
	//})
	//fmt.Println("dialer", dc)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connet to the proxy:", err)
		os.Exit(1)
	}
	httpTransport := &http.Transport{}
	httpTransport.Dial = dialer.Dial
	httpClient := &http.Client{Transport: httpTransport}
	if resp, err := httpClient.Get("https://www.baidu.com"); err != nil {
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
	}
}
