/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:19:29
*/
package main

import (
	"fmt"
	"net/http"
	"strings"
)

/*
NewRequest中method参数对req.Method值的影响
*/
func main() {
	_, err := http.NewRequest("bad method", "http://foo.com/", nil)
	if err == nil { //返回没有输出则说明"bad method"是错误的请求方法，err != nil
		fmt.Println("expected error from NewRequest with invalid method")
	}
	fmt.Println(err) //net/http: invalid method "bad method"

	req, err := http.NewRequest("GET", "http://foo.example/", nil)
	if err != nil { //当你使用的是正确的请求方法时，就不会出现错误
		fmt.Println(err)
	}
	req.Method = "bad method" //将请求方法改成错误的"bad method"

	_, err = http.DefaultClient.Do(req)                                 //然后发送该请求，然后会返回HTTP response和error
	if err == nil || !strings.Contains(err.Error(), "invalid method") { //没有返回，则说明返回的err !=  nil或err中包含字符串"invalid method"
		fmt.Printf("Transport error = %v; want invalid method\n", err)
	}
	fmt.Println(err) //bad method http://foo.example/: net/http: invalid method "bad method"

	req, err = http.NewRequest("", "http://foo.com/", nil)
	fmt.Println(req) //&{GET http://foo.com/ HTTP/1.1 1 1 map[] <nil> <nil> 0 [] false foo.com map[] map[] <nil> map[]   <nil> <nil> <nil> <nil>}
	if err != nil {  //没返回说明err == nil,说明请求方法可以为空
		fmt.Printf("NewRequest(empty method) = %v; want nil\n", err)
	} else if req.Method != "GET" { //当请求方法为空时，会默认使用的是"GET方法"
		fmt.Printf("NewRequest(empty method) has method %q; want GET\n", req.Method)
	}
}
