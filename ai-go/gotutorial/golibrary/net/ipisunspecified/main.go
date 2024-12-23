/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 15:39:49
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6Unspecified := net.ParseIP("::")
	ipv6Specified := net.ParseIP("fe00::")
	ipv4Unspecified := net.ParseIP("0.0.0.0")
	ipv4Specified := net.ParseIP("8.8.8.8")

	fmt.Println(ipv6Unspecified.IsUnspecified())
	fmt.Println(ipv6Specified.IsUnspecified())
	fmt.Println(ipv4Unspecified.IsUnspecified())
	fmt.Println(ipv4Specified.IsUnspecified())

}
