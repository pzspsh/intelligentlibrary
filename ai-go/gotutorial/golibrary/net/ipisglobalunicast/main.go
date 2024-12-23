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
	ipv6Global := net.ParseIP("2000::")
	ipv6UniqLocal := net.ParseIP("2000::")
	ipv6Multi := net.ParseIP("FF00::")

	ipv4Private := net.ParseIP("10.255.0.0")
	ipv4Public := net.ParseIP("8.8.8.8")
	ipv4Broadcast := net.ParseIP("255.255.255.255")

	fmt.Println(ipv6Global.IsGlobalUnicast())
	fmt.Println(ipv6UniqLocal.IsGlobalUnicast())
	fmt.Println(ipv6Multi.IsGlobalUnicast())

	fmt.Println(ipv4Private.IsGlobalUnicast())
	fmt.Println(ipv4Public.IsGlobalUnicast())
	fmt.Println(ipv4Broadcast.IsGlobalUnicast())

}
