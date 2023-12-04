/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:49:58
*/
package main

import (
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	// http.NotFound(w, r)//用于设置404问题
	// http.Error(w, "404 page not found", 404) //状态码需和描述相符

	http.ServeFile(w, r, "1.txt") //将1.txt中内容在w中显示.
	cookie := &http.Cookie{
		Name:  http.CanonicalHeaderKey("uid-test"), //Name值为Uid-Test
		Value: "1234",
	}
	r.AddCookie(cookie)
	fmt.Println(r.Cookie("uid-test")) //<nil> http: named cookie not present
	fmt.Println(r.Cookie("Uid-Test")) //Uid-Test=1234 <nil>
	fmt.Println(r.Cookies())          //[Uid-Test=1234]

}

func main() {
	stat := http.StatusText(200)
	fmt.Println(stat) //状态码200对应的状态OK

	stringtype := http.DetectContentType([]byte("test"))
	fmt.Println(stringtype) //text/plain; charset=utf-8

	http.HandleFunc("/test", Test)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println(err)
	}
}
