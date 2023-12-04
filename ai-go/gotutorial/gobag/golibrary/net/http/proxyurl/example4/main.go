/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 14:52:04
*/
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/armon/go-socks5"
	"github.com/elazarl/goproxy"
	"golang.org/x/net/proxy"
)

type HttpProxy struct {
}

func (HttpProxy) IsProxy(proxyIp string, proxyPort int) (isProxy bool, err error) {
	proxyUrl := fmt.Sprintf("http://%s:%d", proxyIp, proxyPort)
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		return false, err
	}
	netTransport := &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}
	client := &http.Client{
		Timeout:   time.Second * 20, //设置连接超时时间
		Transport: netTransport,
	}
	return CheckProxy(client)
}

const TestUrl = "http://www.baidu.com"

type Proxyer interface {
	IsProxy(proxyIp string, proxyPort int) (isProxy bool, err error)
}

func CheckProxy(client *http.Client) (isProxy bool, err error) {
	res, err := client.Get(TestUrl)
	if err != nil {
		return false, err
	} else {
		defer res.Body.Close()
		if res.StatusCode == 200 {
			body, err := io.ReadAll(res.Body)
			if err == nil && strings.Contains(string(body), "baidu") {
				return true, nil
			} else {
				return false, err
			}
		} else {
			return false, nil
		}
	}
}

type SocksProxy struct {
}

func (SocksProxy) IsProxy(proxyIp string, proxyPort int) (isProxy bool, err error) {
	proxyAddr := fmt.Sprintf("%s:%d", proxyIp, proxyPort)
	dialSocksProxy, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
	if err != nil {
		return false, nil
	}
	netTransport := &http.Transport{DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
		c, e := dialSocksProxy.Dial(network, addr)
		return c, e
	}}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	return CheckProxy(client)
}

func HttpProxyRun() {
	go func() {
		proxy := goproxy.NewProxyHttpServer()
		proxy.Verbose = true
		log.Fatal(http.ListenAndServe(":8080", proxy))
	}()
	time.Sleep(time.Second * 2)
	isProxy, err := HttpProxy{}.IsProxy("127.0.0.1", 8080)
	if !isProxy {
		fmt.Println("should be a proxy", err)
	}
}

func SocksProxyRun() {
	go func() {
		conf := &socks5.Config{}
		server, err := socks5.New(conf)
		if err != nil {
			panic(err)
		}
		if err := server.ListenAndServe("tcp", "127.0.0.1:8002"); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second * 2)
	isProxy, err := SocksProxy{}.IsProxy("127.0.0.1", 8002)
	if !isProxy {
		fmt.Println("should be a proxy", err)
	}
}

/*
使用一个map来存放我们扫出来的代理，最后记录下扫描用时
*/
var (
	Threads = make(chan int, 500) //控制在500以内
)

func RunMain() {
	var start = 8000
	var end = 10000
	var proxyIp = "183.30.204.91"
	var now = time.Now().Unix()
	var Map = make(map[int]bool)
	var waitGroup = sync.WaitGroup{}
	for port := start; port < end; port++ {
		Threads <- 1
		waitGroup.Add(1)
		go func(port int) {
			defer waitGroup.Add(-1)
			isProxy, err := HttpProxy{}.IsProxy(proxyIp, port)
			if isProxy {
				fmt.Printf("%s:%d\n", proxyIp, port)
				Map[port] = true
			}
			if err != nil {
				fmt.Println(err)
			}
			<-Threads
		}(port)
	}
	waitGroup.Wait()
	fmt.Printf("用时%d秒\n", time.Now().Unix()-now)
	fmt.Println(Map)
}

func main() {
	// HttpProxyRun()
	// SocksProxyRun()
	// RunMain()
}
