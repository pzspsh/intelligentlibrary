/*
@File   : main.go
@Author : pan
@Time   : 2024-05-06 17:31:56
*/
package main

import (
	"bufio"
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
)

func handleHTTPSRequest(clientConn net.Conn) {
	defer clientConn.Close()

	// 从客户端读取TLS握手
	clientTLSConn := tls.Server(clientConn, &tls.Config{Certificates: []tls.Certificate{mustLoadCertificate()}})
	/*
	 if err != nil {
	        log.Printf("TLS handshake error from client: %v", err)
	        return
	    }
	*/
	defer clientTLSConn.Close()

	// 从TLS连接读取HTTP请求
	// clientReader := tls.NewReader(clientTLSConn)
	clientReader := bufio.NewReader(clientTLSConn)
	req, err := http.ReadRequest(clientReader)
	if err != nil {
		log.Printf("Failed to read HTTP request from client: %v", err)
		return
	}

	// 创建到目标服务器的连接
	targetConn, err := net.Dial("tcp", req.Host)
	if err != nil {
		log.Printf("Failed to connect to target: %v", err)
		return
	}
	defer targetConn.Close()

	// 将请求转发到目标服务器
	targetTLSConn := tls.Client(targetConn, &tls.Config{InsecureSkipVerify: true}) // 在生产环境中，不应该跳过证书验证

	// targetWriter := tls.NewWriter(targetTLSConn)
	targetWriter := bufio.NewWriter(targetTLSConn)
	err = req.Write(targetWriter)
	if err != nil {
		log.Printf("Failed to write HTTP request to target: %v", err)
		return
	}
	err = targetWriter.Flush()
	if err != nil {
		log.Printf("Failed to flush HTTP request to target: %v", err)
		return
	}

	// 双向复制数据：从目标服务器到客户端，以及从客户端到目标服务器
	go io.Copy(clientTLSConn, targetTLSConn)
	io.Copy(targetTLSConn, clientTLSConn)
}

func mustLoadCertificate() tls.Certificate {
	// 在这里加载你的证书和私钥
	// 这通常涉及到从文件系统中读取它们
	// 这里只是一个占位符，你需要用实际的证书和私钥替换它
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("Failed to load certificate: %v", err)
	}
	return cert
}

func main() {
	listener, err := net.Listen("tcp", ":8443") // 监听HTTPS代理的端口
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	log.Println("HTTPS proxy server listening on :8443")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleHTTPSRequest(conn) // 为每个连接启动一个处理协程
	}
}
