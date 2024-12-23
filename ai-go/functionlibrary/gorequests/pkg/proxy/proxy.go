/*
@File   : proxy.go
@Author : pan
@Time   : 2023-08-24 11:47:49
*/
package proxy

import (
	"net"
)

type DialFunc func(addr string) (net.Conn, error)
