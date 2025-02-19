/*
@File   : main.go
@Author : pan
@Time   : 2025-02-19 17:32:23
*/
package main

import (
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	// 目标服务器地址
	serverAddr := "www.baidu.com:443" // 替换为你的目标服务器地址

	// 建立 TLS 连接
	conn, err := tls.Dial("tcp", serverAddr, &tls.Config{
		InsecureSkipVerify: true, // 跳过证书验证（仅用于测试）
	})
	if err != nil {
		log.Fatalf("无法建立 TLS 连接: %v", err)
	}
	defer conn.Close()

	// 获取连接状态
	state := conn.ConnectionState()

	// 遍历证书链
	for i, cert := range state.PeerCertificates {
		fmt.Printf("=== 证书 %d ===\n", i+1)

		// 打印证书基本信息
		fmt.Println("主题 (Subject):", cert.Subject)
		fmt.Println("颁发者 (Issuer):", cert.Issuer)
		fmt.Println("有效期 (NotBefore):", cert.NotBefore)
		fmt.Println("有效期 (NotAfter):", cert.NotAfter)

		// 解析扩展信息
		fmt.Println("扩展信息 (Extensions):")
		for _, ext := range cert.Extensions {
			fmt.Printf("  - ID: %s\n", ext.Id)
			fmt.Printf("    Critical: %v\n", ext.Critical)
			fmt.Printf("    Value: %s\n", hex.EncodeToString(ext.Value))
			fmt.Println("    --------------------")
		}

		// 打印特定扩展信息
		if len(cert.CRLDistributionPoints) > 0 {
			fmt.Println("CRL 分发点 (crlDistributionPoints):")
			for _, url := range cert.CRLDistributionPoints {
				fmt.Println("  -", url)
			}
		}

		if len(cert.OCSPServer) > 0 {
			fmt.Println("OCSP 服务器 (OCSPServer):")
			for _, url := range cert.OCSPServer {
				fmt.Println("  -", url)
			}
		}

		if len(cert.IssuingCertificateURL) > 0 {
			fmt.Println("CA 证书下载地址 (IssuingCertificateURL):")
			for _, url := range cert.IssuingCertificateURL {
				fmt.Println("  -", url)
			}
		}

		fmt.Println()
	}
}
