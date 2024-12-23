/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 14:13:20
*/
package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	// "git.tencent.com/tke/p2p/pkg/util"
	"github.com/elazarl/goproxy"
	"github.com/gorilla/mux"
	"k8s.io/klog"
)

func main() {
	go func() {
		log.Println("Starting httpServer")
		router := mux.NewRouter().SkipClean(true)
		proxy := goproxy.NewProxyHttpServer()
		proxy.Verbose = true
		proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			req.URL.Host = req.Host
			req.URL.Scheme = "http"
			proxy.ServeHTTP(w, req)
		})
		proxy.OnRequest(goproxy.ReqHostIs("test.openssl.com:1213")).HijackConnect(func(req *http.Request, client net.Conn, _ *goproxy.ProxyCtx) {
			var err error

			log.Printf("getHijhack: %+v", req.URL)
			defer func() {
				if err != nil {
					klog.Errorf("Transfer HTTP CONNECT request failed: %+v, %v", req, err)
					if _, writeErr := client.Write([]byte("HTTP/1.1 500 Cannot reach destination\r\n\r\n")); err != nil {
						klog.Errorf("Write CONNECT failing header failed: %v", writeErr)
					}
				}
				if closeErr := client.Close(); closeErr != nil {
					klog.Errorf("Close client connection failed: %v", closeErr)
				}
			}()

			log.Println("before connectDial")
			remote, err := connectDial(proxy, "tcp", "127.0.0.1:1213")
			if remote != nil {
				log.Printf("==============> remote: %+v>%+v\n", remote.LocalAddr(), remote.RemoteAddr())
			}
			if err != nil {
				return
			}

			bufferedRemote := bufio.NewReadWriter(bufio.NewReader(remote), bufio.NewWriter(remote))
			bufferedClient := bufio.NewReadWriter(bufio.NewReader(client), bufio.NewWriter(client))

			errCh := make(chan error, 1)
			go func() {
				defer close(errCh)
				if _, reverseErr := io.Copy(bufferedRemote, bufferedClient); reverseErr != nil {
					klog.Errorf("Transfer remote to client failed: %v", reverseErr)
					errCh <- reverseErr
				}
			}()

			if _, transferErr := io.Copy(bufferedClient, bufferedRemote); transferErr != nil {
				klog.Errorf("Transfer client to remote failed: %v", transferErr)
				err = transferErr
			}

			if reverseErr := <-errCh; reverseErr != nil {
				err = reverseErr
			}
		})
		router.HandleFunc("/http", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("1--------------------->http: /http >>>>>> req.URL: %+v", r.URL)
			cnt, err := w.Write([]byte(fmt.Sprintf("http: /http return response of req: %+v", r)))
			log.Printf("/http write: cnt: %v, err: %v", cnt, err)
		})
		router.HandleFunc("/https", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("2--------------------->http: /https >>>>>>req.URL: %+v", r.URL)
			cnt, err := w.Write([]byte(fmt.Sprintf("http: /https return response of req: %+v", r)))
			log.Printf("/http write: cnt: %v, err: %v", cnt, err)
			//proxy.ServeHTTP(w, r)
		})
		router.NotFoundHandler = proxy
		if err := http.ListenAndServe(":1212", router); err != nil {
			log.Printf("httpServer err: %+v", err)
		}
	}()
	go func() {
		log.Println("Starting httpsServer")
		router := mux.NewRouter().SkipClean(true)
		proxy := goproxy.NewProxyHttpServer()
		proxy.Verbose = true
		proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			req.URL.Host = req.Host
			req.URL.Scheme = "https"
			proxy.ServeHTTP(w, req)
		})
		if tr, err := TLSTransport("/home/ao/Documents/certs/review/server.crt"); err == nil {
			proxy.Tr = tr
		}
		router.HandleFunc("/https", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("3--------------------->https: req: %+v", r)
			cnt, err := w.Write([]byte(fmt.Sprintf("https: /https return response of req: %+v", r)))
			log.Printf("/http write: cnt: %v, err: %v", cnt, err)
		})
		if err := http.ListenAndServeTLS(":1213", "/home/ao/Documents/certs/review/server.crt", "/home/ao/Documents/certs/review/server.key", router); err != nil {
			log.Printf("httsServer err: %+v", err)
		}
	}()
	select {}
}

func dial(proxy *goproxy.ProxyHttpServer, network, addr string) (c net.Conn, err error) {
	if proxy.Tr.DialContext != nil {
		return proxy.Tr.DialContext(context.Background(), network, addr)
	}
	return net.Dial(network, addr)
}

func connectDial(proxy *goproxy.ProxyHttpServer, network, addr string) (c net.Conn, err error) {
	if proxy.ConnectDial == nil {
		return dial(proxy, network, addr)
	}
	return proxy.ConnectDial(network, addr)
}

func TLSTransport(caFile string) (*http.Transport, error) {
	tr := &http.Transport{TLSClientConfig: &tls.Config{}, Proxy: http.ProxyFromEnvironment}
	if len(caFile) == 0 {
		tr.TLSClientConfig.InsecureSkipVerify = true
		return tr, nil
	}

	ca, err := os.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("read CA file failed: %v", err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(ca)

	tr.TLSClientConfig.RootCAs = pool

	return tr, nil
}
