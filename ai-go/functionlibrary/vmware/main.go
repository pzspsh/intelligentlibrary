package main

import (
	"context"
	"fmt"
	"net/url"

	govmomi "function/vmware/vmware"
)

type Vmware struct {
	Host     string
	Port     int64
	Proxy    string
	Username string // 用户名
	Password string // 密码
	Cmd      string
}

func (v *Vmware) Vmwareconnect() {
	u := &url.URL{
		Scheme: "https",
		Host:   v.Host,
		Path:   "/sdk",
	}
	u.Parse(v.Proxy)
	u.User = url.UserPassword(v.Username, v.Password)
	_, err := govmomi.NewClient(context.Background(), u, v.Proxy, true)
	if err != nil {
		fmt.Println("连接失败")
	} else {
		fmt.Println("连接成功")
	}
}

func main() {

}
