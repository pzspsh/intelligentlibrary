/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:22:08
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
)

var readRequestErrorTests = []struct {
	in  string
	err string

	header http.Header
}{
	0: {"GET / HTTP/1.1\r\nheader:foo\r\n\r\n", "", http.Header{"Header": {"foo"}}},
	1: {"GET / HTTP/1.1\r\nheader:foo\r\n", io.ErrUnexpectedEOF.Error(), nil},
	2: {"", io.EOF.Error(), nil},
	3: {
		in:  "HEAD / HTTP/1.1\r\nContent-Length:4\r\n\r\n",
		err: "http: method cannot contain a Content-Length",
	},
	4: {
		in:     "HEAD / HTTP/1.1\r\n\r\n",
		header: http.Header{},
	},

	// Multiple Content-Length values should either be
	// deduplicated if same or reject otherwise
	// See Issue 16490.
	5: {
		in:  "POST / HTTP/1.1\r\nContent-Length: 10\r\nContent-Length: 0\r\n\r\nGopher hey\r\n",
		err: "cannot contain multiple Content-Length headers",
	},
	6: {
		in:  "POST / HTTP/1.1\r\nContent-Length: 10\r\nContent-Length: 6\r\n\r\nGopher\r\n",
		err: "cannot contain multiple Content-Length headers",
	},
	7: {
		in:     "PUT / HTTP/1.1\r\nContent-Length: 6 \r\nContent-Length: 6\r\nContent-Length:6\r\n\r\nGopher\r\n",
		err:    "",
		header: http.Header{"Content-Length": {"6"}},
	},
	8: {
		in:  "PUT / HTTP/1.1\r\nContent-Length: 1\r\nContent-Length: 6 \r\n\r\n",
		err: "cannot contain multiple Content-Length headers",
	},
	9: {
		in:  "POST / HTTP/1.1\r\nContent-Length:\r\nContent-Length: 3\r\n\r\n",
		err: "cannot contain multiple Content-Length headers",
	},
	10: {
		in:     "HEAD / HTTP/1.1\r\nContent-Length:0\r\nContent-Length: 0\r\n\r\n",
		header: http.Header{"Content-Length": {"0"}},
	},
}

func main() {
	for i, tt := range readRequestErrorTests {
		req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(tt.in)))
		if err == nil { //从返回可以看出，只有0，4，7，10返回的err是nil,即能够成功解析出一个HTTP请求
			fmt.Println(i, " : ", req)
			if tt.err != "" {
				fmt.Printf("#%d: got nil err; want %q\n", i, tt.err)
			}

			if !reflect.DeepEqual(tt.header, req.Header) { //如果发现两者不同
				fmt.Printf("#%d: gotHeader: %q wantHeader: %q\n", i, req.Header, tt.header)
			}
			continue
		}

		if tt.err == "" || !strings.Contains(err.Error(), tt.err) { //如果tt.err != "" 或者 返回的err中包含tt.err的内容，则不会输出下面的字符串
			fmt.Printf("%d: got error = %v; want %v\n", i, err, tt.err)
		}
		fmt.Println(i, "when err is not nil : ", err)
	}
}
