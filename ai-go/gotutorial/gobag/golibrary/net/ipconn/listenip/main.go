/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 17:07:08
*/
package main

import (
	"fmt"
	"net"

	"golang.org/x/net/icmp"
)

/*
ICMP协议是一个“错误侦测与回报机制”，其目的是检测网路的连线状况﹐确保连线的准确性﹐就是我们经常使用的Ping命令。我们在Go中实践下，
来拦截Ping命令产生的数据流量：
*/
func main() {
	netaddr, _ := net.ResolveIPAddr("ip4", "172.17.0.3")
	conn, _ := net.ListenIP("ip4:icmp", netaddr)
	for {
		buf := make([]byte, 1024)
		n, addr, _ := conn.ReadFrom(buf)
		msg, _ := icmp.ParseMessage(1, buf[0:n])
		fmt.Println(n, addr, msg.Type, msg.Code, msg.Checksum)
	}
}
