package main

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

func TrimSpaceRun() {
	IPRegex := regexp.MustCompile(`^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`)
	IPCRegex := regexp.MustCompile(`^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)-((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`)
	IPCSRegex := regexp.MustCompile(`^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)-(25[0-5]|2[0-4]\d|[01]?\d\d?)$`)
	IPSRegex := regexp.MustCompile(`^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)/(3[0-2]|[12]?\d)$`)
	// ipcidr := "192.168.1.0/24,192.168.0.0/16,192.0.0.0/8,192.168.1.0/30, 192.168.10.100-192.168.10.200\n 192.168.1.10-100"
	ipcidr := `192.168.1.0/32,

	192.168.0.0/16,192.0.0.0/8,

	192.168.1.0/30, 192.168.10.100-192.168.10.200, 192.168.1.10-100, 192.168.1.10`
	// fmt.Println(strings.Split(ipcidr, ","))
	for _, v := range strings.Split(ipcidr, ",") {
		// for _, v1 := range strings.Split(v, "\n") {
		// ip := strings.TrimSpace(v1)
		ip := strings.TrimSpace(v)
		if len(ip) > 0 && (IPRegex.MatchString(ip) || IPSRegex.MatchString(ip) || IPCRegex.MatchString(ip) || IPCSRegex.MatchString(ip)) {
			fmt.Println(ip)
			// }
		}
	}
	str := " \t\n Hello, Gophers \n\t\r\n"
	result := strings.TrimSpace(str)
	fmt.Println(str)
	fmt.Println(result)
}

// 解析完整 IP 范围（如 192.168.10.100-192.168.10.200）
func parseFullIPRange(ipRange string) ([]string, error) {
	parts := strings.Split(ipRange, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("无效的 IP 范围格式: %s", ipRange)
	}

	startIP := net.ParseIP(parts[0])
	endIP := net.ParseIP(parts[1])

	if startIP == nil || endIP == nil {
		return nil, fmt.Errorf("无效的 IP 地址: %s 或 %s", parts[0], parts[1])
	}

	// 确保是 IPv4 地址
	startIP = startIP.To4()
	endIP = endIP.To4()

	if startIP == nil || endIP == nil {
		return nil, fmt.Errorf("仅支持 IPv4 地址")
	}

	// 将 IP 地址转换为整数
	startInt := ipToInt(startIP)
	endInt := ipToInt(endIP)

	if startInt > endInt {
		return nil, fmt.Errorf("起始 IP 地址不能大于结束 IP 地址")
	}

	// 生成范围内的所有 IP 地址
	var ips []string
	for i := startInt; i <= endInt; i++ {
		ips = append(ips, intToIP(i).String())
	}

	return ips, nil
}

// 解析简化 IP 范围（如 192.168.1.10-200）
func parseSimplifiedIPRange(ipRange string) ([]string, error) {
	parts := strings.Split(ipRange, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("无效的 IP 范围格式: %s", ipRange)
	}

	baseParts := strings.Split(parts[0], ".")
	if len(baseParts) != 4 {
		return nil, fmt.Errorf("无效的 IP 地址: %s", parts[0])
	}

	startLastOctet, err := strconv.Atoi(baseParts[3])
	if err != nil {
		return nil, fmt.Errorf("无效的起始 IP 地址最后一段: %s", baseParts[3])
	}

	endLastOctet, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("无效的结束 IP 地址最后一段: %s", parts[1])
	}

	if startLastOctet > endLastOctet {
		return nil, fmt.Errorf("起始 IP 地址的最后一段不能大于结束 IP 地址的最后一段")
	}

	// 构造范围内的所有 IP 地址
	var ips []string
	for i := startLastOctet; i <= endLastOctet; i++ {
		ip := fmt.Sprintf("%s.%s.%s.%d", baseParts[0], baseParts[1], baseParts[2], i)
		ips = append(ips, ip)
	}

	return ips, nil
}

// 将 IPv4 地址转换为整数
func ipToInt(ip net.IP) uint32 {
	ip = ip.To4()
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
}

// 将整数转换为 IPv4 地址
func intToIP(n uint32) net.IP {
	return net.IPv4(byte(n>>24), byte(n>>16), byte(n>>8), byte(n))
}

func IPRangeRun() {
	// 测试完整 IP 范围
	fullRange := "192.168.10.100-192.168.10.200"
	fullIPs, err := parseFullIPRange(fullRange)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("完整 IP 范围 (%s):\n%v\n", fullRange, fullIPs)
	}

	// 测试简化 IP 范围
	simplifiedRange := "192.168.1.10-200"
	simplifiedIPs, err := parseSimplifiedIPRange(simplifiedRange)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("简化 IP 范围 (%s):\n%v\n", simplifiedRange, simplifiedIPs)
	}
}

func main() {
	TrimSpaceRun()
	// IPRangeRun()
}
