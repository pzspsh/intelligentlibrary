/*
@File   : util.go
@Author : pan
@Time   : 2023-08-24 15:26:30
*/
package http

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"

	"gorequests/pkg/client"
	readerutil "gorequests/pkg/utils/reader"
	urlutil "gorequests/pkg/utils/url"
)

type ContextOverride string

const (
	RETRY_MAX ContextOverride = "retry-max"
)

// Discard is an helper function that discards the response body and closes the underlying connection
func Discard(req *Request, resp *http.Response, RespReadLimit int64) {
	_, err := io.Copy(io.Discard, io.LimitReader(resp.Body, RespReadLimit))
	if err != nil {
		req.Metrics.DrainErrors++
	}
	resp.Body.Close()
}

// getLength returns length of a Reader efficiently
func getLength(x io.Reader) (int64, error) {
	len, err := io.Copy(io.Discard, x)
	return len, err
}

func getReusableBodyandContentLength(rawBody interface{}) (*readerutil.ReusableReadCloser, int64, error) {

	var bodyReader *readerutil.ReusableReadCloser
	var contentLength int64

	if rawBody != nil {
		switch body := rawBody.(type) {
		// If they gave us a function already, great! Use it.
		case readerutil.ReusableReadCloser:
			bodyReader = &body
		case *readerutil.ReusableReadCloser:
			bodyReader = body
		// If they gave us a reader function read it and get reusablereader
		case func() (io.Reader, error):
			tmp, err := body()
			if err != nil {
				return nil, 0, err
			}
			bodyReader, err = readerutil.NewReusableReadCloser(tmp)
			if err != nil {
				return nil, 0, err
			}
		// If ReusableReadCloser is not given try to create new from it
		// if not possible return error
		default:
			var err error
			bodyReader, err = readerutil.NewReusableReadCloser(body)
			if err != nil {
				return nil, 0, err
			}
		}
	}

	if bodyReader != nil {
		var err error
		contentLength, err = getLength(bodyReader)
		if err != nil {
			return nil, 0, err
		}
	}

	return bodyReader, contentLength, nil
}

// StatusError is a HTTP status error object
type StatusError struct {
	client.Status
}

func (s *StatusError) Error() string {
	return s.Status.String()
}

type readCloser struct {
	io.Reader
	io.Closer
}

func toRequest(method string, path string, query []string, headers map[string][]string, body io.Reader, options *Options) *client.Request {
	if len(options.CustomRawBytes) > 0 {
		return &client.Request{RawBytes: options.CustomRawBytes}
	}
	reqHeaders := toHeaders(headers)
	if len(options.CustomHeaders) > 0 {
		reqHeaders = options.CustomHeaders
	}

	return &client.Request{
		Method:  method,
		Path:    path,
		Query:   query,
		Version: client.HTTP_1_1,
		Headers: reqHeaders,
		Body:    body,
	}
}
func toHTTPResponse(conn Conn, resp *client.Response) (*http.Response, error) {
	rheaders := fromHeaders(resp.Headers)
	r := http.Response{
		ProtoMinor:    resp.Version.Minor,
		ProtoMajor:    resp.Version.Major,
		Status:        resp.Status.String(),
		StatusCode:    resp.Status.Code,
		Header:        rheaders,
		ContentLength: resp.ContentLength(),
	}

	var err error
	rbody := resp.Body
	if headerValue(rheaders, "Content-Encoding") == "gzip" {
		rbody, err = gzip.NewReader(rbody)
		if err != nil {
			return nil, err
		}
	}
	rc := &readCloser{rbody, conn}

	r.Body = rc

	return &r, nil
}

func toHeaders(h map[string][]string) []client.Header {
	var r []client.Header
	for k, v := range h {
		for _, v := range v {
			r = append(r, client.Header{Key: k, Value: v})
		}
	}
	return r
}

func fromHeaders(h []client.Header) map[string][]string {
	if h == nil {
		return nil
	}
	var r = make(map[string][]string)
	for _, hh := range h {
		r[hh.Key] = append(r[hh.Key], hh.Value)
	}
	return r
}

func headerValue(headers map[string][]string, key string) string {
	return strings.Join(headers[key], " ")
}

func firstErr(err1, err2 error) error {
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

// DumpRequestRaw to string
func DumpRequestRaw(method, url, uripath string, headers map[string][]string, body io.Reader, options *Options) ([]byte, error) {
	if len(options.CustomRawBytes) > 0 {
		return options.CustomRawBytes, nil
	}
	if headers == nil {
		headers = make(map[string][]string)
	}
	u, err := urlutil.ParseURL(url, true)
	if err != nil {
		return nil, err
	}

	// Handle only if host header is missing
	_, hasHostHeader := headers["Host"]
	if !hasHostHeader {
		host := u.Host
		headers["Host"] = []string{host}
	}

	// standard path
	path := u.Path
	if path == "" {
		path = "/"
	}
	if !u.Params.IsEmpty() {
		path += "?" + u.Params.Encode()
	}
	// override if custom one is specified
	if uripath != "" {
		path = uripath
	}

	req := toRequest(method, path, nil, headers, body, options)
	b := strings.Builder{}

	q := strings.Join(req.Query, "&")
	if len(q) > 0 {
		q = "?" + q
	}

	b.WriteString(fmt.Sprintf("%s %s%s %s"+client.NewLine, req.Method, req.Path, q, req.Version.String()))

	for _, header := range req.Headers {
		if header.Value != "" {
			b.WriteString(fmt.Sprintf("%s: %s"+client.NewLine, header.Key, header.Value))
		} else {
			b.WriteString(fmt.Sprintf("%s"+client.NewLine, header.Key))
		}
	}

	l := req.ContentLength()
	if req.AutomaticContentLength && l >= 0 {
		b.WriteString(fmt.Sprintf("Content-Length: %d"+client.NewLine, l))
	}

	b.WriteString(client.NewLine)

	if req.Body != nil {
		var buf bytes.Buffer
		tee := io.TeeReader(req.Body, &buf)
		body, err := io.ReadAll(tee)
		if err != nil {
			return nil, err
		}
		b.Write(body)
	}

	return []byte(strings.ReplaceAll(b.String(), "\n", client.NewLine)), nil
}
