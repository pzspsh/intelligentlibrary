/*
@File   : main.go
@Author : pan
@Time   : 2024-12-09 10:31:29
*/
package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 创建http客户端
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 10,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 10 * time.Second,
	}
	// 发送https请求
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("请求失败：", err)
		return
	}
	defer resp.Body.Close()

	// 获取TLS连接信息
	if resp.TLS == nil {
		fmt.Println("TLS连接信息为空")
		return
	}

	// 获取证书信息
	cert := resp.TLS.PeerCertificates[0]
	// 输出证书信息
	fmt.Println("证书信息：")
	fmt.Println("  序列号：", cert.SerialNumber)
	fmt.Println("  主题：", cert.Subject)
	fmt.Println("  颁发者：", cert.Issuer)
	fmt.Println("  有效期：", cert.NotBefore, "至", cert.NotAfter)
	fmt.Println("  签名算法：", cert.SignatureAlgorithm)
	fmt.Println("  公钥算法：", cert.PublicKeyAlgorithm)
	fmt.Println("  公钥：", cert.PublicKey)
	fmt.Println("  扩展：", cert.Extensions)
	fmt.Println("  签名：", cert.Signature)
	fmt.Println("  key id：", cert.AuthorityKeyId)
	fmt.Println("  版本号：", cert.Version)
}
