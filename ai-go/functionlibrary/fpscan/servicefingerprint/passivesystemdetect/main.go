/*
@File   : main.go
@Author : pan
@Time   : 2024-06-11 11:57:29
*/
package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/ipv4"
)

// 假设的操作系统特征
const (
	osAWindowMinSize = 1000
	osAWindowMaxSize = 2000
	osBTimestampMin  = 1234567890
	osBTimestampMax  = 9876543210
)

// parseTCPOptions 解析TCP选项并返回它们
func parseTCPOptions(rawOptions []byte) map[byte][]byte {
	options := make(map[byte][]byte)
	optLen := 0
	for i := 0; i < len(rawOptions); i += optLen {
		opt := rawOptions[i]
		if opt == 0 { // NOP选项，跳过
			optLen = 1
			continue
		}
		optLen = int(rawOptions[i+1])
		if optLen < 2 {
			log.Fatalf("Invalid option length: %d", optLen)
		}
		options[opt] = rawOptions[i+2 : i+optLen]
	}
	return options
}

// analyzePacket 分析数据包并尝试识别操作系统
func analyzePacket(ipHeader ipv4.Header, tcpHeader []byte) string {
	windowSize := int(binary.BigEndian.Uint16(tcpHeader[14:16]))
	timestamp := 0
	fmt.Println(ipHeader)
	tcpOptions := parseTCPOptions(tcpHeader[20:])
	timestampOption, ok := tcpOptions[0x08]
	if ok && len(timestampOption) >= 8 {
		timestamp = int(binary.BigEndian.Uint32(timestampOption[4:8]))
	}

	// 根据提取的字段进行操作系统识别
	if windowSize >= osAWindowMinSize && windowSize <= osAWindowMaxSize {
		return "可能是操作系统A"
	}
	if timestamp >= osBTimestampMin && timestamp <= osBTimestampMax {
		return "可能是操作系统B"
	}
	return "未知的操作系统"
}

// capturePackets 使用原始套接字捕获TCP SYN和SYN-ACK包
func capturePackets() {
	// 创建一个原始套接字，用于捕获IPv4上的TCP流量
	// 创建一个原始套接字，用于捕获IPv4上的TCP流量
	conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		log.Fatalf("Failed to create raw socket: %v", err)
	}
	defer conn.Close()
	buffer := make([]byte, 65535)
	for {
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Printf("Error reading from socket: %v", err)
			continue
		}

		// 解析IP头部
		ipHeader, err := ipv4.ParseHeader(buffer[:n])
		if err != nil {
			log.Printf("Error parsing IP header: %v", err)
			continue
		}

		// TCP头部在IP数据部分开始的位置
		tcpHeader := buffer[ipHeader.Len:]

		// 分析数据包并打印结果
		os := analyzePacket(*ipHeader, tcpHeader)
		fmt.Printf("捕获到来自 %s 的数据包，可能是 %s\n", addr, os)
	}
}

func main() {
	// 创建一个原始套接字，用于捕获IPv4上的TCP流量
	capturePackets()
}
