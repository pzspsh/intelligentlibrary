/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 16:36:44
*/
package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Serverslice struct {
	Servers []Server `json:"servers"`
}

type Server struct {
	ServerName string `json:"servername"`
	ServerIP   string `json:"serverip,omitempty"`
}

func main() {
	var s Serverslice
	//func append(slice []Type, elems ...Type) []Type
	s.Servers = append(s.Servers, Server{ServerName: "Beijing", ServerIP: "10.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Xi'an", ServerIP: "10.0.0.2"})
	//slice里面嵌套结构体[{},{}] 遍历出来的是slice里面包含json串
	ss := []Server{{"Beijing", "172.0.0.1"}, {"Shanghai", "172.0.0.2"}}
	b, err := json.Marshal(s)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(b))

	bb, err := json.MarshalIndent(ss, "", "  ")
	if err == nil {
		fmt.Println(string(bb))
	}

}
