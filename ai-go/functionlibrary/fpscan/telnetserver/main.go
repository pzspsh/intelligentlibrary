/*
@File   : main.go
@Author : pan
@Time   : 2024-06-07 15:15:54
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func detectTelnetService(target string) (string, error) {
	// 设置连接超时
	conn, err := net.DialTimeout("tcp", target+":23", 5*time.Second)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	// 设置读取超时
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	// 创建读取器
	reader := bufio.NewReader(conn)
	var banner string
	for {
		// 尝试读取一行数据，直到超时或读取到换行符
		line, err := reader.ReadString('\n')
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				// 读取超时，返回当前收集到的banner（如果有）
				return banner, nil
			} else {
				// 其他错误，返回错误
				return "", err
			}
		}
		banner += line

		// 检查是否读取到了完整的banner（这取决于你的具体需求）
		// 或者是否达到了某个读取次数的限制
		// 如果满足条件，则跳出循环
		if strings.Contains(banner, "some ending condition") {
			break
		}
	}
	return banner, nil
}

func main() {
	target := "10.0.35.74" // 替换为目标IP地址和端口（格式：IP:PORT）
	banner, err := detectTelnetService(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to Telnet service: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Telnet banner: %s\n", banner)
}
