/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 12:25:23
*/
package main

import (
	"fmt"
	"net/http"
)

type logWrites struct {
	dst *[]string
}

// 实现Write函数说明logWrites实现了io.Writer接口
func (l logWrites) Write(p []byte) (n int, err error) {
	*l.dst = append(*l.dst, string(p))
	return len(p), nil
}

func main() {
	got1 := []string{}
	got2 := []string{}
	req, _ := http.NewRequest("GET", "http://foo.com/", nil)
	fmt.Println(req)
	req.Write(logWrites{&got1})      //logWrites{&got}得到的是一个io.Writer对象作为req.Write的参数，这样就会自动调用func (l logWrites) Write(p []byte)，将req写入got中
	req.WriteProxy(logWrites{&got2}) //logWrites{&got}得到的是一个io.Writer对象作为req.Write的参数，这样就会自动调用func (l logWrites) Write(p []byte)，将req写入got中
	fmt.Println(got1)
	fmt.Println(got2)
}
