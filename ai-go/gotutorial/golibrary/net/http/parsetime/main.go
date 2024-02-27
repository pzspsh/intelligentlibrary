/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:08:10
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

var parseTimeTests = []struct {
	h   http.Header
	err bool
}{
	{http.Header{"Date": {""}}, true},
	{http.Header{"Date": {"invalid"}}, true},
	{http.Header{"Date": {"1994-11-06T08:49:37Z00:00"}}, true},
	{http.Header{"Date": {"Sun, 06 Nov 1994 08:49:37 GMT"}}, false},
	{http.Header{"Date": {"Sunday, 06-Nov-94 08:49:37 GMT"}}, false},
	{http.Header{"Date": {"Sun Nov 6 08:49:37 1994"}}, false},
}

func main() {
	expect := time.Date(1994, 11, 6, 8, 49, 37, 0, time.UTC)
	fmt.Println(expect) //1994-11-06 08:49:37 +0000 UTC
	for i, test := range parseTimeTests {
		d, err := http.ParseTime(test.h.Get("Date"))
		fmt.Println(d)
		if err != nil {
			fmt.Println(i, err)
			if !test.err { //test.err为false才进这里
				fmt.Printf("#%d:\n got err: %v", i, err)
			}
			continue //有错的进入这后继续下一个循环，不往下执行
		}
		if test.err { //test.err为true，所以该例子中这里不会进入
			fmt.Printf("#%d:\n should err", i)
			continue
		}
		if !expect.Equal(d) { //说明后三个例子的结果和expect是相同的，所以没有报错
			fmt.Printf("#%d:\n got: %v\nwant: %v", i, d, expect)
		}
	}
}
