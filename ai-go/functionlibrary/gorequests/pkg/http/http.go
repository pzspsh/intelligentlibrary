package http

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"time"
)

var DefaultClient1 = Client{
	dialer:  new(dialer),
	Options: DefaultOptions,
}

func DefaultHostSprayingTransport() *http.Transport {
	transport := DefaultReusePooledTransport()
	transport.DisableKeepAlives = true
	transport.MaxIdleConnsPerHost = -1
	return transport
}

func DefaultReusePooledTransport() *http.Transport {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:           100,
		IdleConnTimeout:        90 * time.Second,
		TLSHandshakeTimeout:    10 * time.Second,
		ExpectContinueTimeout:  1 * time.Second,
		MaxIdleConnsPerHost:    100,
		MaxResponseHeaderBytes: 4096, // net/http default is 10Mb
		TLSClientConfig: &tls.Config{
			Renegotiation:      tls.RenegotiateOnceAsClient,
			InsecureSkipVerify: true,
			MinVersion:         tls.VersionTLS10,
		},
	}
	return transport
}

func DefaultClient() *http.Client {
	return &http.Client{
		Transport: DefaultHostSprayingTransport(),
	}
}

func DefaultPooledClient() *http.Client {
	return &http.Client{
		Transport: DefaultReusePooledTransport(),
	}
}

// Get makes a GET request to a given URL
func Get(url string) (*http.Response, error) {
	return DefaultClient1.Get(url)
}

// Post makes a POST request to a given URL
func Post(url string, mimetype string, r io.Reader) (*http.Response, error) {
	return DefaultClient1.Post(url, mimetype, r)
}

// Do sends a http request and returns a response
func Do(req *http.Request) (*http.Response, error) {
	return DefaultClient1.Do(req)
}

// Dor sends a retryablehttp request and returns a response
func Dor(req *Request) (*http.Response, error) {
	return DefaultClient1.Dor(req)
}

// DoRaw does a raw request with some configuration
func DoRaw(method, url, uripath string, headers map[string][]string, body io.Reader) (*http.Response, error) {
	return DefaultClient1.DoRaw(method, url, uripath, headers, body)
}

// DoRawWithOptions does a raw request with some configuration
func DoRawWithOptions(method, url, uripath string, headers map[string][]string, body io.Reader, options *Options) (*http.Response, error) {
	return DefaultClient1.DoRawWithOptions(method, url, uripath, headers, body, options)
}
