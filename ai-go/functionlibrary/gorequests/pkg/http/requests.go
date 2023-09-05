/*
@File   : requests.go
@Author : pan
@Time   : 2023-08-24 14:04:40
*/
package http

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"
	"net/url"
	"os"

	readerutil "gorequests/pkg/utils/reader"
	urlutil "gorequests/pkg/utils/url"
)

var PreferHTTP bool

type Request struct {
	*http.Request
	*urlutil.URL
	Metrics Metrics
	Auth    *Auth
}

// Metrics contains the metrics about each request
type Metrics struct {
	// Failures is the number of failed requests
	Failures int
	// Retries is the number of retries for the request
	Retries int
	// DrainErrors is number of errors occured in draining response body
	DrainErrors int
}

// Auth specific information
type Auth struct {
	Type     AuthType
	Username string
	Password string
}

type AuthType uint8

const (
	DigestAuth AuthType = iota
)

type RequestLogHook func(*http.Request, int)

type ResponseLogHook func(*http.Response)

type ErrorHandler func(resp *http.Response, err error, numTries int) (*http.Response, error)

func (r *Request) WithContext(ctx context.Context) *Request {
	r.Request = r.Request.WithContext(ctx)
	return r
}

func (r *Request) BodyBytes() ([]byte, error) {
	if r.Request.Body == nil {
		return nil, nil
	}
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Update request URL with new changes of parameters if any
func (r *Request) Update() {
	r.URL.Update()
	updateScheme(r.URL.URL)
}

// SetURL updates request url (i.e http.Request.URL) with given url
func (r *Request) SetURL(u *urlutil.URL) {
	r.URL = u
	r.Request.URL = u.URL
	r.Update()
}

// Clones and returns new Request
func (r *Request) Clone(ctx context.Context) *Request {
	r.Update()
	ux := r.URL.Clone()
	req := r.Request.Clone(ctx)
	req.URL = ux.URL
	ux.Update()
	var auth *Auth
	if r.hasAuth() {
		auth = &Auth{
			Type:     r.Auth.Type,
			Username: r.Auth.Username,
			Password: r.Auth.Password,
		}
	}
	return &Request{
		Request: req,
		URL:     ux,
		Metrics: Metrics{}, // Metrics shouldn't be cloned
		Auth:    auth,
	}
}

// Dump returns request dump in bytes
func (r *Request) Dump() ([]byte, error) {
	resplen := int64(0)
	dumpbody := true
	clone := r.Clone(context.TODO())
	if clone.Body != nil {
		resplen, _ = getLength(clone.Body)
	}
	if resplen == 0 {
		dumpbody = false
		clone.ContentLength = 0
		clone.Body = nil
		delete(clone.Header, "Content-length")
	}
	dumpBytes, err := httputil.DumpRequestOut(clone.Request, dumpbody)
	if err != nil {
		return nil, err
	}
	return dumpBytes, nil
}

// hasAuth checks if request has any username/password
func (request *Request) hasAuth() bool {
	return request.Auth != nil
}

// FromRequest wraps an http.Request in a retryablehttp.Request
func FromRequest(r *http.Request) (*Request, error) {
	req := Request{
		Request: r,
		Metrics: Metrics{},
		Auth:    nil,
	}

	if r.URL != nil {
		urlx, err := urlutil.Parse(r.URL.String())
		if err != nil {
			return nil, err
		}
		req.URL = urlx
	}

	if r.Body != nil {
		body, err := readerutil.NewReusableReadCloser(r.Body)
		if err != nil {
			return nil, err
		}
		r.Body = body
		req.ContentLength, err = getLength(body)
		if err != nil {
			return nil, err
		}
	}

	return &req, nil
}

// FromRequestWithTrace wraps an http.Request in a retryablehttp.Request with trace enabled
func FromRequestWithTrace(r *http.Request) (*Request, error) {
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Fprintf(os.Stderr, "Got connection\tReused: %v\tWas Idle: %v\tIdle Time: %v\n", connInfo.Reused, connInfo.WasIdle, connInfo.IdleTime)
		},
		ConnectStart: func(network, addr string) {
			fmt.Fprintf(os.Stderr, "Dial start\tnetwork: %s\taddress: %s\n", network, addr)
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Fprintf(os.Stderr, "Dial done\tnetwork: %s\taddress: %s\terr: %v\n", network, addr, err)
		},
		GotFirstResponseByte: func() {
			fmt.Fprintf(os.Stderr, "Got response's first byte\n")
		},
		WroteHeaders: func() {
			fmt.Fprintf(os.Stderr, "Wrote request headers\n")
		},
		WroteRequest: func(wr httptrace.WroteRequestInfo) {
			fmt.Fprintf(os.Stderr, "Wrote request, err: %v\n", wr.Err)
		},
	}

	r = r.WithContext(httptrace.WithClientTrace(r.Context(), trace))

	return FromRequest(r)
}

// NewRequest creates a new wrapped request.
func NewRequestFromURL(method string, urlx *urlutil.URL, body interface{}) (*Request, error) {
	return NewRequestFromURLWithContext(context.Background(), method, urlx, body)
}

// NewRequestWithContext creates a new wrapped request with context
func NewRequestFromURLWithContext(ctx context.Context, method string, urlx *urlutil.URL, body interface{}) (*Request, error) {
	bodyReader, contentLength, err := getReusableBodyandContentLength(body)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, method, "https://"+urlx.Host, nil)
	if err != nil {
		return nil, err
	}
	urlx.Update()
	httpReq.URL = urlx.URL
	updateScheme(httpReq.URL)
	// content-length and body should be assigned only
	// if request has body
	if bodyReader != nil {
		httpReq.ContentLength = contentLength
		httpReq.Body = bodyReader
	}

	return &Request{httpReq, urlx, Metrics{}, nil}, nil
}

// NewRequest creates a new wrapped request
func NewRequest(method, url string, body interface{}) (*Request, error) {
	urlx, err := urlutil.Parse(url)
	if err != nil {
		return nil, err
	}
	return NewRequestFromURL(method, urlx, body)
}

// NewRequest creates a new wrapped request with given context
func NewRequestWithContext(ctx context.Context, method, url string, body interface{}) (*Request, error) {
	urlx, err := urlutil.Parse(url)
	if err != nil {
		return nil, err
	}
	return NewRequestFromURLWithContext(ctx, method, urlx, body)
}

func updateScheme(u *url.URL) {
	if u.Host != "" && u.Scheme == "" {
		if PreferHTTP {
			u.Scheme = "http"
		} else {
			u.Scheme = "https"
		}
	}
}
