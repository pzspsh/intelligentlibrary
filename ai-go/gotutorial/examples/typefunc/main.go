/*
@File   : main.go
@Author : pan
@Time   : 2024-12-24 15:59:24
*/
package main

// type func 类型函数的使用教程

import (
	"fmt"
	"net"
	"time"
)

// IsOpen 检查指定主机和端口是否在指定超时时间内开放
func IsOpen(host string, port int, timeout time.Duration) (bool, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", host, port), timeout)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	return true, nil
}

// IsOpenScan 是一个函数类型，与 IsOpen 函数签名相同
type IsOpenScan func(host string, port int, timeout time.Duration) (bool, error)

// ScanChan 是一个容量为 1000 的通道，用于存储 IsOpenScan 类型的函数
var ScanChan = make(chan IsOpenScan, 1000)

func main() {
	// 将 IsOpen 函数发送到 ScanChan 通道
	ScanChan <- IsOpen

	// 从 ScanChan 通道接收一个函数，并调用它来检查端口是否开放
	go func() {
		scanFunc := <-ScanChan
		host := "www.example.com"
		port := 80
		timeout := 2 * time.Second

		open, err := scanFunc(host, port, timeout)
		if err != nil {
			fmt.Printf("Error scanning %s:%d - %v\n", host, port, err)
		} else if open {
			fmt.Printf("Port %d on host %s is open\n", port, host)
		} else {
			fmt.Printf("Port %d on host %s is closed\n", port, host)
		}
	}()

	// 为了确保主 goroutine 不提前退出，等待一段时间
	time.Sleep(3 * time.Second)
}
