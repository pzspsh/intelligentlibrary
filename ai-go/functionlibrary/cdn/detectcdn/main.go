/*
@File   : main.go
@Author : pan
@Time   : 2024-11-27 15:07:25
*/
package main

import (
	"fmt"
	"net"
)

// 假设的CDN IP地址列表（实际使用时需要更新和维护）
var cdnIPs = map[string]bool{
	"192.0.2.1":   true, // 示例CDN IP
	"203.0.113.1": true, // 示例CDN IP
	// ... 添加更多CDN IP地址
}

// isCDNIP 检查IP地址是否在CDN IP列表中
func isCDNIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	// 对于IPv4，直接比较字符串；对于IPv6，你可能需要更复杂的比较逻辑
	return cdnIPs[parsedIP.String()]
}

func main() {
	// 测试用例
	ips := []string{
		"192.0.2.1",   // CDN IP
		"192.168.1.1", // 用户IP
		"203.0.113.1", // CDN IP
		"8.8.8.8",     // Google DNS（不是CDN，但用于示例）
	}

	for _, ip := range ips {
		if isCDNIP(ip) {
			fmt.Printf("%s is a CDN IP address.\n", ip)
		} else {
			fmt.Printf("%s is NOT a CDN IP address.\n", ip)
		}
	}
}

/*
这个示例代码中的cdnIPs映射是硬编码的，并且非常不完整。在实际应用中，你需要维护一个更完整且定期更新的CDN IP地址列表。
此外，对于IPv6地址和更复杂的CDN配置（如使用多个IP范围或动态分配的IP），你可能需要实现更复杂的检查逻辑。

最终，验证IP地址是否属于CDN是一个复杂的问题，没有简单的解决方案。你可能需要结合多种方法来提高准确性。
*/
