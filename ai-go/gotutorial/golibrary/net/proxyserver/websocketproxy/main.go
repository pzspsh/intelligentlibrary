/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 15:36:43
*/
package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

const (
	WsScheme  = "ws"
	WssScheme = "wss"
	BufSize   = 1024 * 32
)

var ErrFormatAddr = errors.New("remote websockets addr format error")

type WebsocketProxy struct {
	// ws, wss
	scheme string
	// The target address: host:port
	remoteAddr string
	// path
	defaultPath string
	tlsc        *tls.Config
	logger      *log.Logger
	// Send handshake before callback
	beforeHandshake func(r *http.Request) error
}

type Options func(wp *WebsocketProxy)

// You must carry a port number，ws://ip:80/ssss, wss://ip:443/aaaa
// ex: ws://ip:port/ajaxchattest
func NewProxy(addr string, beforeCallback func(r *http.Request) error, options ...Options) (*WebsocketProxy, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, ErrFormatAddr
	}
	host, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		return nil, ErrFormatAddr
	}
	if u.Scheme != WsScheme && u.Scheme != WssScheme {
		return nil, ErrFormatAddr
	}
	wp := &WebsocketProxy{
		scheme:          u.Scheme,
		remoteAddr:      fmt.Sprintf("%s:%s", host, port),
		defaultPath:     u.Path,
		beforeHandshake: beforeCallback,
		logger:          log.New(os.Stderr, "", log.LstdFlags),
	}
	if u.Scheme == WssScheme {
		wp.tlsc = &tls.Config{InsecureSkipVerify: true}
	}
	for op := range options {
		options[op](wp)
	}
	return wp, nil
}

func (wp *WebsocketProxy) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	wp.Proxy(writer, request)
}

func (wp *WebsocketProxy) Proxy(writer http.ResponseWriter, request *http.Request) {
	if strings.ToLower(request.Header.Get("Connection")) != "upgrade" ||
		strings.ToLower(request.Header.Get("Upgrade")) != "websocket" {
		_, _ = writer.Write([]byte(`Must be a websocket request`))
		return
	}
	hijacker, ok := writer.(http.Hijacker)
	if !ok {
		return
	}
	conn, _, err := hijacker.Hijack()
	if err != nil {
		return
	}
	defer conn.Close()
	req := request.Clone(request.Context())
	req.URL.Path, req.URL.RawPath, req.RequestURI = wp.defaultPath, wp.defaultPath, wp.defaultPath
	req.Host = wp.remoteAddr
	if wp.beforeHandshake != nil {
		// Add headers, permission authentication + masquerade sources
		err = wp.beforeHandshake(req)
		if err != nil {
			_, _ = writer.Write([]byte(err.Error()))
			return
		}
	}
	var remoteConn net.Conn
	switch wp.scheme {
	case WsScheme:
		remoteConn, err = net.Dial("tcp", wp.remoteAddr)
	case WssScheme:
		remoteConn, err = tls.Dial("tcp", wp.remoteAddr, wp.tlsc)
	}
	if err != nil {
		_, _ = writer.Write([]byte(err.Error()))
		return
	}
	defer remoteConn.Close()
	err = req.Write(remoteConn)
	if err != nil {
		wp.logger.Println("remote write err:", err)
		return
	}
	errChan := make(chan error, 2)
	copyConn := func(a, b net.Conn) {
		buf := ByteSliceGet(BufSize)
		defer ByteSlicePut(buf)
		_, err := io.CopyBuffer(a, b, buf)
		errChan <- err
	}
	go copyConn(conn, remoteConn) // response
	go copyConn(remoteConn, conn) // request
	select {
	case err = <-errChan:
		if err != nil {
			log.Println(err)
		}
	}
}

func SetTLSConfig(tlsc *tls.Config) Options {
	return func(wp *WebsocketProxy) {
		wp.tlsc = tlsc
	}
}

func SetLogger(l *log.Logger) Options {
	return func(wp *WebsocketProxy) {
		if l != nil {
			wp.logger = l
		}
	}
}

var (
	byteSlicePool = sync.Pool{
		New: func() interface{} {
			return []byte{}
		},
	}
	byteSliceChan = make(chan []byte, 10)
)

func ByteSliceGet(length int) (data []byte) {
	select {
	case data = <-byteSliceChan:
	default:
		data = byteSlicePool.Get().([]byte)[:0]
	}

	if cap(data) < length {
		data = make([]byte, length)
	} else {
		data = data[:length]
	}

	return data
}

func ByteSlicePut(data []byte) {
	select {
	case byteSliceChan <- data:
	default:
		byteSlicePool.Put(data) // nolint:staticcheck
	}
}

func auth(r *http.Request) error {
	// Permission to verify
	r.Header.Set("Cookie", "----")
	// Source of disguise
	r.Header.Set("Origin", "http://82.157.123.54:9010")
	return nil
}

func main() {
	tlsc := tls.Config{InsecureSkipVerify: true}
		wp, err := NewProxy("ws://www.baidu.com:80/ajaxchattest", auth, SetTLSConfig(&tlsc))
		if err != nil {
			fmt.Println(err)
		}
	http.HandleFunc("/wsproxy", wp.Proxy)
	http.ListenAndServe(":9696", nil)
}
