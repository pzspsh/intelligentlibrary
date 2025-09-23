package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./osdetect <IP> <PORT>")
		return
	}

	ip := os.Args[1]
	port, _ := strconv.Atoi(os.Args[2])

	result := detectOS(ip, port)
	fmt.Printf("[%s:%d] 操作系统可能是: %s\n", ip, port, result)
}

func detectOS(ip string, port int) string {
	// 综合检测逻辑
	if os := tcpFingerprint(ip, port); os != "" {
		return os
	}
	if os := icmpTTLDetection(ip); os != "" {
		return os
	}
	return servicePortAnalysis(port)
}

func tcpFingerprint(ip string, port int) string {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%v", ip, port), 3*time.Second)
	if err != nil {
		return ""
	}
	defer conn.Close()

	// 发送探测数据
	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	buf := make([]byte, 256)
	n, _ := conn.Read(buf)
	response := string(buf[:n])

	// 分析响应特征
	switch {
	case regexp.MustCompile(`Linux|Ubuntu|CentOS`).MatchString(response):
		return "Linux"
	case regexp.MustCompile(`Windows|IIS`).MatchString(response):
		return "Windows"
	case regexp.MustCompile(`Cisco`).MatchString(response):
		return "Cisco IOS"
	}
	return ""
}

func icmpTTLDetection(ip string) string {
	cmd := exec.Command("ping", "-c", "1", ip)
	output, _ := cmd.CombinedOutput()

	ttlRegex := regexp.MustCompile(`ttl=(\d+)`)
	matches := ttlRegex.FindStringSubmatch(string(output))
	if len(matches) < 2 {
		return ""
	}

	ttl, _ := strconv.Atoi(matches[1])
	switch {
	case ttl <= 64:
		return "Linux/Unix"
	case ttl <= 128:
		return "Windows"
	default:
		return "Network Device"
	}
}

func servicePortAnalysis(port int) string {
	portMap := map[int]string{
		3389: "Windows (RDP)",
		22:   "Linux (SSH)",
		161:  "Network Device (SNMP)",
		443:  "跨平台 (HTTPS)",
	}
	return portMap[port]
}
