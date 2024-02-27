/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:11:37
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	cname, srvs, err := net.LookupSRV("xmpp-server", "tcp", "golang.org")
	if err != nil {
		panic(err)
	}

	fmt.Printf("\ncname: %s \n\n", cname)

	for _, srv := range srvs {
		fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}
}
