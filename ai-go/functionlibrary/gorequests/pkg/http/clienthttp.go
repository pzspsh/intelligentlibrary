/*
@File   : clienthttp.go
@Author : pan
@Time   : 2023-08-24 16:12:02
*/
package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	dac "function/gorequests/pkg/digestauth"
)

type ClientHttp struct {
	HttpClient      *http.Client
	HttpClient2     *http.Client
	requestCounter  uint32
	options         OptionsClient
	RequestLogHook  RequestLogHook
	ResponseLogHook ResponseLogHook
	ErrorHandler    ErrorHandler
	CheckRetry      CheckRetry
	Backoff         Backoff
}

func PassthroughErrorHandler(resp *http.Response, err error, _ int) (*http.Response, error) {
	return resp, err
}

func (c *ClientHttp) Get(url string) (*http.Response, error) {
	req, err := NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *ClientHttp) Head(url string) (*http.Response, error) {
	req, err := NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *ClientHttp) Post(url string, bodyType string, body interface{}) (*http.Response, error) {
	req, err := NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	return c.Do(req)
}

func (c *ClientHttp) PostForm(url string, data url.Values) (*http.Response, error) {
	return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func (c *ClientHttp) Do(req *Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	// Create a main context that will be used as the main timeout
	mainCtx, cancel := context.WithTimeout(context.Background(), c.options.Timeout)
	defer cancel()

	retryMax := c.options.RetryMax
	if ctxRetryMax := req.Context().Value(RETRY_MAX); ctxRetryMax != nil {
		if maxRetriesParsed, ok := ctxRetryMax.(int); ok {
			retryMax = maxRetriesParsed
		}
	}
	for i := 0; ; i++ {
		if c.RequestLogHook != nil {
			c.RequestLogHook(req.Request, i)
		}

		if req.hasAuth() && req.Auth.Type == DigestAuth {
			digestTransport := dac.NewTransport(req.Auth.Username, req.Auth.Password)
			digestTransport.HTTPClient = c.HttpClient
			resp, err = digestTransport.RoundTrip(req.Request)
		} else {
			// Attempt the request with standard behavior
			resp, err = c.HttpClient.Do(req.Request)
		}
		// Check if we should continue with retries.
		checkOK, checkErr := c.CheckRetry(req.Context(), resp, err)

		// if err is equal to missing minor protocol version retry with http/2
		if err != nil && strings.Contains(err.Error(), "net/http: HTTP/1.x transport connection broken: malformed HTTP version \"HTTP/2\"") {
			resp, err = c.HttpClient2.Do(req.Request)
			checkOK, checkErr = c.CheckRetry(req.Context(), resp, err)
		}
		if err != nil {
			// Increment the failure counter as the request failed
			req.Metrics.Failures++
		} else {
			if c.ResponseLogHook != nil {
				// Call the response logger function if provided.
				c.ResponseLogHook(resp)
			}
		}
		// Now decide if we should continue.
		if !checkOK {
			if checkErr != nil {
				err = checkErr
			}
			c.closeIdleConnections()
			return resp, err
		}
		remain := retryMax - i
		if remain <= 0 {
			break
		}
		// Increment the retries counter as we are going to do one more retry
		req.Metrics.Retries++
		// We're going to retry, consume any response to reuse the connection.
		if err == nil && resp != nil {
			c.drainBody(req, resp)
		}
		wait := c.Backoff(c.options.RetryWaitMin, c.options.RetryWaitMax, i, resp)
	selectstatement:
		select {
		case <-mainCtx.Done():
			break selectstatement
		case <-req.Context().Done():
			c.closeIdleConnections()
			return nil, req.Context().Err()
		case <-time.After(wait):
		}
	}

	if c.ErrorHandler != nil {
		c.closeIdleConnections()
		return c.ErrorHandler(resp, err, retryMax+1)
	}
	if resp != nil {
		resp.Body.Close()
	}
	c.closeIdleConnections()
	return nil, fmt.Errorf("%s %s giving up after %d attempts: %w", req.Method, req.URL, retryMax+1, err)
}

// Try to read the response body so we can reuse this connection.
func (c *ClientHttp) drainBody(req *Request, resp *http.Response) {
	_, err := io.Copy(io.Discard, io.LimitReader(resp.Body, c.options.RespReadLimit))
	if err != nil {
		req.Metrics.DrainErrors++
	}
	resp.Body.Close()
}

const closeConnectionsCounter = 100

func (c *ClientHttp) closeIdleConnections() {
	if c.options.KillIdleConn {
		requestCounter := atomic.LoadUint32(&c.requestCounter)
		if requestCounter < closeConnectionsCounter {
			atomic.AddUint32(&c.requestCounter, 1)
		} else {
			atomic.StoreUint32(&c.requestCounter, 0)
			c.HttpClient.CloseIdleConnections()
		}
	}
}
