/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 17:25:40
*/
package main

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"regexp"

	"github.com/elazarl/goproxy"
)

func orPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest(goproxy.ReqHostMatches(regexp.MustCompile("^.*baidu.com$"))).
		HandleConnect(goproxy.AlwaysReject)
	proxy.OnRequest(goproxy.ReqHostMatches(regexp.MustCompile("^.*$"))).
		HandleConnect(goproxy.AlwaysMitm)
	// enable curl -p for all hosts on port 80
	proxy.OnRequest(goproxy.ReqHostMatches(regexp.MustCompile("^.*:80$"))).
		HijackConnect(func(req *http.Request, client net.Conn, ctx *goproxy.ProxyCtx) {
			defer func() {
				if e := recover(); e != nil {
					ctx.Logf("error connecting to remote: %v", e)
					client.Write([]byte("HTTP/1.1 500 Cannot reach destination\r\n\r\n"))
				}
				client.Close()
			}()
			clientBuf := bufio.NewReadWriter(bufio.NewReader(client), bufio.NewWriter(client))
			remote, err := net.Dial("tcp", req.URL.Host)
			orPanic(err)
			remoteBuf := bufio.NewReadWriter(bufio.NewReader(remote), bufio.NewWriter(remote))
			for {
				req, err := http.ReadRequest(clientBuf.Reader)
				orPanic(err)
				orPanic(req.Write(remoteBuf))
				orPanic(remoteBuf.Flush())
				resp, err := http.ReadResponse(remoteBuf.Reader, req)
				orPanic(err)
				orPanic(resp.Write(clientBuf.Writer))
				orPanic(clientBuf.Flush())
			}
		})
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":8080", proxy))
}

// https://vimsky.com/examples/detail/golang-ex-bufio---NewReadWriter-function.html
