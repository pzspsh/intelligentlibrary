/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:30:58
*/
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func main() {
	req, _ := http.NewRequest("POST", "http://www.google.com/search?q=foo&q=bar&both=x&prio=1&orphan=nope&empty=not",
		strings.NewReader("z=post&both=y&prio=2&=nokey&orphan;empty=&"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	if q := req.FormValue("q"); q != "foo" {
		fmt.Printf(`req.FormValue("q") = %q, want "foo"`, q)
	}
	//因为上面的req.FormValue方法会隐式解析，所以下面能够得到值
	fmt.Println(req)
	fmt.Println(req.Form)
	fmt.Println(req.PostForm)

	if z := req.FormValue("z"); z != "post" {
		fmt.Printf(`req.FormValue("z") = %q, want "post"`, z)
	}
	if bq, found := req.PostForm["q"]; found { //PostForm 中没有"q"
		fmt.Printf(`req.PostForm["q"] = %q, want no entry in map`, bq)
	}
	if bz := req.PostFormValue("z"); bz != "post" {
		fmt.Printf(`req.PostFormValue("z") = %q, want "post"`, bz)
	}
	if qs := req.Form["q"]; !reflect.DeepEqual(qs, []string{"foo", "bar"}) {
		fmt.Printf(`req.Form["q"] = %q, want ["foo", "bar"]`, qs)
	}
	if both := req.Form["both"]; !reflect.DeepEqual(both, []string{"y", "x"}) {
		fmt.Printf(`req.Form["both"] = %q, want ["y", "x"]`, both)
	}
	if prio := req.FormValue("prio"); prio != "2" {
		fmt.Printf(`req.FormValue("prio") = %q, want "2" (from body)`, prio)
	}
	if orphan := req.Form["orphan"]; !reflect.DeepEqual(orphan, []string{"", "nope"}) {
		fmt.Printf(`req.FormValue("orphan") = %q, want "" (from body)`, orphan)
	}
	if empty := req.Form["empty"]; !reflect.DeepEqual(empty, []string{"", "not"}) {
		fmt.Printf(`req.FormValue("empty") = %q, want "" (from body)`, empty)
	}
	if nokey := req.Form[""]; !reflect.DeepEqual(nokey, []string{"nokey"}) {
		fmt.Printf(`req.FormValue("nokey") = %q, want "nokey" (from body)`, nokey)
	}
}
