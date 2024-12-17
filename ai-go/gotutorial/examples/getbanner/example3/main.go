/*
@File   : main.go
@Author : pan
@Time   : 2024-12-16 17:49:47
*/
package main

import (
	"bufio"
	"crypto/tls"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/textproto"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
)

func main() {
	// 获取HTTP服务的Banner信息
	httpBanner := getHTTPBanner("http://www.example.com")
	fmt.Println("HTTP Banner:", httpBanner)

	// 获取HTTPS服务的Banner信息
	httpsBanner := getHTTPSBanner("https://www.example.com")
	fmt.Println("HTTPS Banner:", httpsBanner)

	// 获取FTP服务的Banner信息
	ftpBanner := getFTPBanner("ftp://ftp.example.com")
	fmt.Println("FTP Banner:", ftpBanner)

	// 获取SSH服务的Banner信息
	sshBanner := getSSHBanner("ssh://user:password@example.com:22")
	fmt.Println("SSH Banner:", sshBanner)

	// 获取MySQL服务的Banner信息
	mysqlBanner := getMySQLBanner("user:password@tcp(example.com:3306)/dbname")
	fmt.Println("MySQL Banner:", mysqlBanner)
}

func getHTTPBanner(url string) string {
	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		return "Error sending GET request: " + err.Error()
	}
	defer resp.Body.Close()

	// 读取响应的头部信息
	banner := resp.Header.Get("Server")
	if banner == "" {
		banner = "No banner information found."
	}
	return banner
}

func getHTTPSBanner(url string) string {
	// 创建一个HTTP客户端，并配置TLS
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 忽略证书验证，仅用于测试
			},
		},
	}

	// 发送HTTPS GET请求
	resp, err := client.Get(url)
	if err != nil {
		return "Error sending GET request: " + err.Error()
	}
	defer resp.Body.Close()

	// 读取响应的头部信息
	banner := resp.Header.Get("Server")
	if banner == "" {
		banner = "No banner information found."
	}
	return banner
}

func getFTPBanner(url string) string {
	// 解析FTP URL
	conn, err := net.DialTimeout("tcp", url, 5*time.Second)
	if err != nil {
		return "Error connecting to FTP server: " + err.Error()
	}

	// 创建FTP连接
	ftpConn := textproto.NewConn(conn)
	defer ftpConn.Close()

	// 发送FTP命令获取Banner信息
	ftpConn.Cmd("NOOP")
	ftpConn.StartResponse(0)
	defer ftpConn.EndResponse(0)

	// 读取FTP服务器的响应
	resp, err := ftpConn.ReadLine()
	if err != nil {
		return "Error reading FTP response: " + err.Error()
	}
	return resp
}

func getSSHBanner(url string) string {
	// 解析SSH URL
	config := &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略主机密钥验证，仅用于测试
	}

	// 连接到SSH服务器
	conn, err := ssh.Dial("tcp", url, config)
	if err != nil {
		return "Error connecting to SSH server: " + err.Error()
	}
	defer conn.Close()

	// 读取SSH服务器的Banner信息
	session, err := conn.NewSession()
	if err != nil {
		return "Error creating SSH session: " + err.Error()
	}
	defer session.Close()
	// 执行SSH命令获取Banner信息
	// 获取会话的标准输出管道
	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed to obtain stdout pipe: %s", err)
	}
	reader := bufio.NewReader(stdout)
	// 启动会话（例如，运行一个命令）
	if err := session.Run("your_command"); err != nil {
		log.Fatalf("Failed to run: %s", err)
	}
	// 读取并输出会话的输出
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break // 输出完毕
			}
			log.Fatalf("Error reading from session: %s", err)
		}
		fmt.Printf("Output: %s", line)
	}
	banner, err := reader.ReadString('\n') // 读取到第一个换行符
	if err != nil {
		return "Error reading SSH banner: " + err.Error()
	}
	return string(banner)
}

func getMySQLBanner(url string) string {
	// 解析MySQL URL
	dsn, err := mysql.ParseDSN(url)
	if err != nil {
		return "Error parsing MySQL DSN: " + err.Error()
	}

	// 连接到MySQL服务器
	db, err := sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		return "Error connecting to MySQL server: " + err.Error()
	}
	defer db.Close()

	// 发送SQL查询获取Banner信息
	rows, err := db.Query("SELECT @@version")
	if err != nil {
		return "Error querying MySQL server: " + err.Error()
	}
	defer rows.Close()

	// 读取查询结果
	var version string
	if rows.Next() {
		err := rows.Scan(&version)
		if err != nil {
			return "Error scanning MySQL result: " + err.Error()
		}
	}
	return version
}
