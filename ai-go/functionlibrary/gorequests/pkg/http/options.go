/*
@File   : options.go
@Author : pan
@Time   : 2023-08-24 15:02:01
*/
package http

import (
	"net/http"
	"time"

	"function/gorequests/pkg/client"
	"function/gorequests/pkg/clientpipeline"

	"github.com/projectdiscovery/fastdialer/fastdialer"
	"golang.org/x/net/http2"
)

// Options contains configuration options for rawhttp client
type Options struct {
	Timeout                time.Duration
	FollowRedirects        bool
	MaxRedirects           int
	AutomaticHostHeader    bool
	AutomaticContentLength bool
	CustomHeaders          client.Headers
	ForceReadAllBody       bool // ignores content length and reads all body
	CustomRawBytes         []byte
	Proxy                  string
	ProxyDialTimeout       time.Duration
	SNI                    string
	FastDialer             *fastdialer.Dialer
}

// // DefaultOptions is the default configuration options for the client
var DefaultOptions = &Options{
	Timeout:                30 * time.Second,
	FollowRedirects:        true,
	MaxRedirects:           10,
	AutomaticHostHeader:    true,
	AutomaticContentLength: true,
}

type OptionsClient struct {
	// RetryWaitMin is the minimum time to wait for retry
	RetryWaitMin time.Duration
	// RetryWaitMax is the maximum time to wait for retry
	RetryWaitMax time.Duration
	// Timeout is the maximum time to wait for the request
	Timeout time.Duration
	// RetryMax is the maximum number of retries
	RetryMax int
	// RespReadLimit is the maximum HTTP response size to read for
	// connection being reused.
	RespReadLimit int64
	// Verbose specifies if debug messages should be printed
	Verbose bool
	// KillIdleConn specifies if all keep-alive connections gets killed
	KillIdleConn bool
	// Custom CheckRetry policy
	CheckRetry CheckRetry
	// Custom Backoff policy
	Backoff Backoff
	// NoAdjustTimeout disables automatic adjustment of HTTP request timeout
	NoAdjustTimeout bool
	// Custom http client
	HttpClient *http.Client
}

var DefaultOptionsSpraying = OptionsClient{
	RetryWaitMin:    1 * time.Second,
	RetryWaitMax:    30 * time.Second,
	Timeout:         30 * time.Second,
	RetryMax:        5,
	RespReadLimit:   4096,
	KillIdleConn:    true,
	NoAdjustTimeout: true,
}

var DefaultOptionsSingle = OptionsClient{
	RetryWaitMin:    1 * time.Second,
	RetryWaitMax:    30 * time.Second,
	Timeout:         30 * time.Second,
	RetryMax:        5,
	RespReadLimit:   4096,
	KillIdleConn:    false,
	NoAdjustTimeout: true,
}

// NewClient creates a new Client with default settings.
func NewClient(options OptionsClient) *ClientHttp {
	var httpclient *http.Client
	if options.HttpClient != nil {
		httpclient = options.HttpClient
	} else {
		httpclient = DefaultClient()
	}

	httpclient2 := DefaultClient()
	if err := http2.ConfigureTransport(httpclient2.Transport.(*http.Transport)); err != nil {
		return nil
	}

	var retryPolicy CheckRetry
	var backoff Backoff

	retryPolicy = DefaultRetryPolicy()
	if options.CheckRetry != nil {
		retryPolicy = options.CheckRetry
	}

	backoff = DefaultBackoff()
	if options.Backoff != nil {
		backoff = options.Backoff
	}

	// add timeout to clients
	if options.Timeout > 0 {
		httpclient.Timeout = options.Timeout
		httpclient2.Timeout = options.Timeout
	}

	// if necessary adjusts per-request timeout proportionally to general timeout (30%)
	if options.Timeout > time.Second*15 && options.RetryMax > 1 && !options.NoAdjustTimeout {
		httpclient.Timeout = time.Duration(options.Timeout.Seconds()*0.3) * time.Second
	}

	c := &ClientHttp{
		HttpClient:  httpclient,
		HttpClient2: httpclient2,
		CheckRetry:  retryPolicy,
		Backoff:     backoff,
		options:     options,
	}

	c.setKillIdleConnections()
	return c
}

func NewWithHTTPClient(client *http.Client, options OptionsClient) *ClientHttp {
	options.HttpClient = client
	return NewClient(options)
}

func (c *ClientHttp) setKillIdleConnections() {
	if c.HttpClient != nil || !c.options.KillIdleConn {
		if b, ok := c.HttpClient.Transport.(*http.Transport); ok {
			c.options.KillIdleConn = b.DisableKeepAlives || b.MaxConnsPerHost < 0
		}
	}
}

type PipelineOptions struct {
	Dialer              clientpipeline.DialFunc
	Host                string
	Timeout             time.Duration
	MaxConnections      int
	MaxPendingRequests  int
	AutomaticHostHeader bool
}

// DefaultPipelineOptions is the default options for pipelined http client
var DefaultPipelineOptions = PipelineOptions{
	Timeout:             30 * time.Second,
	MaxConnections:      5,
	MaxPendingRequests:  100,
	AutomaticHostHeader: true,
}
