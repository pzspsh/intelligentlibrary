/*
@File   : main.go
@Author : pan
@Time   : 2024-12-17 11:46:21
*/
package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func IsOpen(host string, port int, timeout time.Duration) (bool, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", host, port), timeout)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	return true, nil
}

func ScanPort(ports []int) {
	var wg sync.WaitGroup
	timeout := 100 * time.Millisecond
	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			IsOpen("", port, timeout)
		}(port)
	}
	wg.Wait()
}

func main() {

}
