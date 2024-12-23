/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:18:11
*/
package main

import (
	"fmt"
	"net/http"
)

type headerOnlyResponseWriter http.Header

// 下面定义这些函数是为了使headerOnlyResponseWriter实现ResponseWriter接口，然后可以作为SetCookie的参数传入
func (ho headerOnlyResponseWriter) Header() http.Header {
	return http.Header(ho)
}

func (ho headerOnlyResponseWriter) Write([]byte) (int, error) {
	panic("NOIMPL")
}

func (ho headerOnlyResponseWriter) WriteHeader(int) {
	panic("NOIMPL")
}

func main() {
	m := make(http.Header)                   //创建一个map[string][]string类型的映射m，headerOnlyResponseWriter(m)即将Header类型的m转成自定义headerOnlyResponseWriter类型
	fmt.Println(m)                           //运行SetCookie()之前为 map[]
	fmt.Println(headerOnlyResponseWriter(m)) //运行SetCookie()之前为 map[]

	//SetCookie在w的头域中添加Set-Cookie头，该HTTP头的值为cookie
	http.SetCookie(headerOnlyResponseWriter(m), &http.Cookie{Name: "cookie-1", Value: "one", Path: "/restricted/"})
	http.SetCookie(headerOnlyResponseWriter(m), &http.Cookie{Name: "cookie-2", Value: "two", MaxAge: 3600})
	fmt.Println(m) //返回：map[Set-Cookie:[cookie-1=one; Path=/restricted/ cookie-2=two; Max-Age=3600]]

	//下面的内容都没有报错，说明得到的值和给出的字符串是相同的
	if l := len(m["Set-Cookie"]); l != 2 {
		fmt.Printf("expected %d cookies, got %d", 2, l)
	}
	if g, e := m["Set-Cookie"][0], "cookie-1=one; Path=/restricted/"; g != e {
		fmt.Printf("cookie #1: want %q, got %q", e, g)
	}
	if g, e := m["Set-Cookie"][1], "cookie-2=two; Max-Age=3600"; g != e {
		fmt.Printf("cookie #2: want %q, got %q", e, g)
	}
}
