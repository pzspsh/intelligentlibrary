/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:35:04
*/
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type respTest struct {
	Raw  string
	Resp http.Response
	Body string
}

func dummyReq(method string) *http.Request {
	return &http.Request{Method: method}
}

var respTests = []respTest{
	// Unchunked response without Content-Length.
	{
		"HTTP/1.0 200 OK\r\n" +
			"Connection: close\r\n" +
			"\r\n" +
			"Body here\n",

		http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.0",
			ProtoMajor: 1,
			ProtoMinor: 0,
			Request:    dummyReq("GET"),
			Header: http.Header{
				"Connection": {"close"}, // TODO(rsc): Delete?
			},
			Close:         true,
			ContentLength: -1,
		},

		"Body here\n",
	},

	// Unchunked HTTP/1.1 response without Content-Length or
	// Connection headers.
	{
		"HTTP/1.1 200 OK\r\n" +
			"\r\n" +
			"Body here\n",

		http.Response{
			Status:        "200 OK",
			StatusCode:    200,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Header:        http.Header{},
			Request:       dummyReq("GET"),
			Close:         true,
			ContentLength: -1,
		},

		"Body here\n",
	},

	// Unchunked HTTP/1.1 204 response without Content-Length.
	{
		"HTTP/1.1 204 No Content\r\n" +
			"\r\n" +
			"Body should not be read!\n",

		http.Response{
			Status:        "204 No Content",
			StatusCode:    204,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Header:        http.Header{},
			Request:       dummyReq("GET"),
			Close:         false,
			ContentLength: 0,
		},

		"",
	},

	// Unchunked response with Content-Length.
	{
		"HTTP/1.0 200 OK\r\n" +
			"Content-Length: 10\r\n" +
			"Connection: close\r\n" +
			"\r\n" +
			"Body here\n",

		http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.0",
			ProtoMajor: 1,
			ProtoMinor: 0,
			Request:    dummyReq("GET"),
			Header: http.Header{
				"Connection":     {"close"},
				"Content-Length": {"10"},
			},
			Close:         true,
			ContentLength: 10,
		},

		"Body here\n",
	},

	// Chunked response without Content-Length.
	{
		"HTTP/1.1 200 OK\r\n" +
			"Transfer-Encoding: chunked\r\n" +
			"\r\n" +
			"0a\r\n" +
			"Body here\n\r\n" +
			"09\r\n" +
			"continued\r\n" +
			"0\r\n" +
			"\r\n",

		http.Response{
			Status:           "200 OK",
			StatusCode:       200,
			Proto:            "HTTP/1.1",
			ProtoMajor:       1,
			ProtoMinor:       1,
			Request:          dummyReq("GET"),
			Header:           http.Header{},
			Close:            false,
			ContentLength:    -1,
			TransferEncoding: []string{"chunked"},
		},

		"Body here\ncontinued",
	},

	// Chunked response with Content-Length.
	{
		"HTTP/1.1 200 OK\r\n" +
			"Transfer-Encoding: chunked\r\n" +
			"Content-Length: 10\r\n" +
			"\r\n" +
			"0a\r\n" +
			"Body here\n\r\n" +
			"0\r\n" +
			"\r\n",

		http.Response{
			Status:           "200 OK",
			StatusCode:       200,
			Proto:            "HTTP/1.1",
			ProtoMajor:       1,
			ProtoMinor:       1,
			Request:          dummyReq("GET"),
			Header:           http.Header{},
			Close:            false,
			ContentLength:    -1,
			TransferEncoding: []string{"chunked"},
		},

		"Body here\n",
	},

	// Chunked response in response to a HEAD request
	{
		"HTTP/1.1 200 OK\r\n" +
			"Transfer-Encoding: chunked\r\n" +
			"\r\n",

		http.Response{
			Status:           "200 OK",
			StatusCode:       200,
			Proto:            "HTTP/1.1",
			ProtoMajor:       1,
			ProtoMinor:       1,
			Request:          dummyReq("HEAD"),
			Header:           http.Header{},
			TransferEncoding: []string{"chunked"},
			Close:            false,
			ContentLength:    -1,
		},

		"",
	},

	// Content-Length in response to a HEAD request
	{
		"HTTP/1.0 200 OK\r\n" +
			"Content-Length: 256\r\n" +
			"\r\n",

		http.Response{
			Status:           "200 OK",
			StatusCode:       200,
			Proto:            "HTTP/1.0",
			ProtoMajor:       1,
			ProtoMinor:       0,
			Request:          dummyReq("HEAD"),
			Header:           http.Header{"Content-Length": {"256"}},
			TransferEncoding: nil,
			Close:            true,
			ContentLength:    256,
		},

		"",
	},

	// Content-Length in response to a HEAD request with HTTP/1.1
	{
		"HTTP/1.1 200 OK\r\n" +
			"Content-Length: 256\r\n" +
			"\r\n",

		http.Response{
			Status:           "200 OK",
			StatusCode:       200,
			Proto:            "HTTP/1.1",
			ProtoMajor:       1,
			ProtoMinor:       1,
			Request:          dummyReq("HEAD"),
			Header:           http.Header{"Content-Length": {"256"}},
			TransferEncoding: nil,
			Close:            false,
			ContentLength:    256,
		},

		"",
	},

	// No Content-Length or Chunked in response to a HEAD request
	{
		"HTTP/1.0 200 OK\r\n" +
			"\r\n",

		http.Response{
			Status:           "200 OK",
			StatusCode:       200,
			Proto:            "HTTP/1.0",
			ProtoMajor:       1,
			ProtoMinor:       0,
			Request:          dummyReq("HEAD"),
			Header:           http.Header{},
			TransferEncoding: nil,
			Close:            true,
			ContentLength:    -1,
		},

		"",
	},

	// explicit Content-Length of 0.
	{
		"HTTP/1.1 200 OK\r\n" +
			"Content-Length: 0\r\n" +
			"\r\n",

		http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Request:    dummyReq("GET"),
			Header: http.Header{
				"Content-Length": {"0"},
			},
			Close:         false,
			ContentLength: 0,
		},

		"",
	},

	// Status line without a Reason-Phrase, but trailing space.
	// (permitted by RFC 7230, section 3.1.2)
	{
		"HTTP/1.0 303 \r\n\r\n",
		http.Response{
			Status:        "303 ",
			StatusCode:    303,
			Proto:         "HTTP/1.0",
			ProtoMajor:    1,
			ProtoMinor:    0,
			Request:       dummyReq("GET"),
			Header:        http.Header{},
			Close:         true,
			ContentLength: -1,
		},

		"",
	},

	// Status line without a Reason-Phrase, and no trailing space.
	// (not permitted by RFC 7230, but we'll accept it anyway)
	{
		"HTTP/1.0 303\r\n\r\n",
		http.Response{
			Status:        "303",
			StatusCode:    303,
			Proto:         "HTTP/1.0",
			ProtoMajor:    1,
			ProtoMinor:    0,
			Request:       dummyReq("GET"),
			Header:        http.Header{},
			Close:         true,
			ContentLength: -1,
		},

		"",
	},

	// golang.org/issue/4767: don't special-case multipart/byteranges responses
	{
		`HTTP/1.1 206 Partial Content
Connection: close
Content-Type: multipart/byteranges; boundary=18a75608c8f47cef

some body`,
		http.Response{
			Status:     "206 Partial Content",
			StatusCode: 206,
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Request:    dummyReq("GET"),
			Header: http.Header{
				"Content-Type": []string{"multipart/byteranges; boundary=18a75608c8f47cef"},
			},
			Close:         true,
			ContentLength: -1,
		},

		"some body",
	},

	// Unchunked response without Content-Length, Request is nil
	{
		"HTTP/1.0 200 OK\r\n" +
			"Connection: close\r\n" +
			"\r\n" +
			"Body here\n",

		http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.0",
			ProtoMajor: 1,
			ProtoMinor: 0,
			Header: http.Header{
				"Connection": {"close"}, // TODO(rsc): Delete?
			},
			Close:         true,
			ContentLength: -1,
		},

		"Body here\n",
	},

	// 206 Partial Content. golang.org/issue/8923
	{
		"HTTP/1.1 206 Partial Content\r\n" +
			"Content-Type: text/plain; charset=utf-8\r\n" +
			"Accept-Ranges: bytes\r\n" +
			"Content-Range: bytes 0-5/1862\r\n" +
			"Content-Length: 6\r\n\r\n" +
			"foobar",

		http.Response{
			Status:     "206 Partial Content",
			StatusCode: 206,
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Request:    dummyReq("GET"),
			Header: http.Header{
				"Accept-Ranges":  []string{"bytes"},
				"Content-Length": []string{"6"},
				"Content-Type":   []string{"text/plain; charset=utf-8"},
				"Content-Range":  []string{"bytes 0-5/1862"},
			},
			ContentLength: 6,
		},

		"foobar",
	},

	// Both keep-alive and close, on the same Connection line. (Issue 8840)
	{
		"HTTP/1.1 200 OK\r\n" +
			"Content-Length: 256\r\n" +
			"Connection: keep-alive, close\r\n" +
			"\r\n",

		http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Request:    dummyReq("HEAD"),
			Header: http.Header{
				"Content-Length": {"256"},
			},
			TransferEncoding: nil,
			Close:            true,
			ContentLength:    256,
		},

		"",
	},

	// Both keep-alive and close, on different Connection lines. (Issue 8840)
	{
		"HTTP/1.1 200 OK\r\n" +
			"Content-Length: 256\r\n" +
			"Connection: keep-alive\r\n" +
			"Connection: close\r\n" +
			"\r\n",

		http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Request:    dummyReq("HEAD"),
			Header: http.Header{
				"Content-Length": {"256"},
			},
			TransferEncoding: nil,
			Close:            true,
			ContentLength:    256,
		},

		"",
	},

	// Issue 12785: HTTP/1.0 response with bogus (to be ignored) Transfer-Encoding.
	// Without a Content-Length.
	{
		"HTTP/1.0 200 OK\r\n" +
			"Transfer-Encoding: bogus\r\n" +
			"\r\n" +
			"Body here\n",

		http.Response{
			Status:        "200 OK",
			StatusCode:    200,
			Proto:         "HTTP/1.0",
			ProtoMajor:    1,
			ProtoMinor:    0,
			Request:       dummyReq("GET"),
			Header:        http.Header{},
			Close:         true,
			ContentLength: -1,
		},

		"Body here\n",
	},

	// Issue 12785: HTTP/1.0 response with bogus (to be ignored) Transfer-Encoding.
	// With a Content-Length.
	{
		"HTTP/1.0 200 OK\r\n" +
			"Transfer-Encoding: bogus\r\n" +
			"Content-Length: 10\r\n" +
			"\r\n" +
			"Body here\n",

		http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.0",
			ProtoMajor: 1,
			ProtoMinor: 0,
			Request:    dummyReq("GET"),
			Header: http.Header{
				"Content-Length": {"10"},
			},
			Close:         true,
			ContentLength: 10,
		},

		"Body here\n",
	},

	{
		"HTTP/1.1 200 OK\r\n" +
			"Content-Encoding: gzip\r\n" +
			"Content-Length: 23\r\n" +
			"Connection: keep-alive\r\n" +
			"Keep-Alive: timeout=7200\r\n\r\n" +
			"\x1f\x8b\b\x00\x00\x00\x00\x00\x00\x00s\xf3\xf7\a\x00\xab'\xd4\x1a\x03\x00\x00\x00",
		http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Request:    dummyReq("GET"),
			Header: http.Header{
				"Content-Length":   {"23"},
				"Content-Encoding": {"gzip"},
				"Connection":       {"keep-alive"},
				"Keep-Alive":       {"timeout=7200"},
			},
			Close:         false,
			ContentLength: 23,
		},
		"\x1f\x8b\b\x00\x00\x00\x00\x00\x00\x00s\xf3\xf7\a\x00\xab'\xd4\x1a\x03\x00\x00\x00",
	},

	// Issue 19989: two spaces between HTTP version and status.
	{
		"HTTP/1.0  401 Unauthorized\r\n" +
			"Content-type: text/html\r\n" +
			"WWW-Authenticate: Basic realm=\"\"\r\n\r\n" +
			"Your Authentication failed.\r\n",
		http.Response{
			Status:     "401 Unauthorized",
			StatusCode: 401,
			Proto:      "HTTP/1.0",
			ProtoMajor: 1,
			ProtoMinor: 0,
			Request:    dummyReq("GET"),
			Header: http.Header{
				"Content-Type":     {"text/html"},
				"Www-Authenticate": {`Basic realm=""`},
			},
			Close:         true,
			ContentLength: -1,
		},
		"Your Authentication failed.\r\n",
	},
}

func main() {
	for i, tt := range respTests {
		resp, err := http.ReadResponse(bufio.NewReader(strings.NewReader(tt.Raw)), tt.Resp.Request)
		if err != nil {
			fmt.Printf("#%d: %v", i, err)
			continue
		}

		fmt.Println(i, resp) //返回得到的response
		fmt.Println("ProtoAtLeast 1.0 : ", resp.ProtoAtLeast(1, 0))
		fmt.Println()

		fmt.Println("write : ")
		err = resp.Write(os.Stdout) //将得到的response写到终端上
		fmt.Println()
		if err != nil {
			fmt.Printf("#%d: %v", i, err)
			continue
		}
	}
}
