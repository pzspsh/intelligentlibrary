/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 00:49:29
*/
package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	hostname := flag.String("hostname", "www.baidu.com", "hostname to test")
	startPort := flag.Int("start-port", 80, "the port on which the scanning starts")
	endPort := flag.Int("end-port", 100, "the port from which the scanning ends")
	timeout := flag.Duration("timeout", time.Millisecond*200, "timeout")
	flag.Parse()

	ports := make([]int, 0)
	//同步机制
	wg := &sync.WaitGroup{}
	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				ports = append(ports, p)
			}
			wg.Done()
		}(port)
	}
	wg.Wait()
	fmt.Printf("opened ports: %v\n", ports)
	/*
		opened ports: [80]
	*/
}

/*
检测端口是否可连接
*/
func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}
