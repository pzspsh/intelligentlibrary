# 端口探活
```go
package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
    "context"
)

var (
	file, _ = os.Create("test.txt")
)

func IsOpen(host string, port int, timeout time.Duration) (bool, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", host, port), timeout)
	if err == nil {
		defer conn.Close()
		return true, err
	}
	return false, nil
}

func PortScan(ip string, ports []int) {
	var wg sync.WaitGroup
	timeout := 200 * time.Millisecond
	threads := make(chan bool, 5000)
	start := time.Now()
	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		threads <- true
		go func(port int, ch chan bool) {
			defer wg.Done()
			isopen, _ := IsOpen(ip, port, timeout)
			if isopen {
				ports = append(ports, port)
			}
			<-ch
		}(port, threads)
	}
	wg.Wait()
	fmt.Println(ports)
	fmt.Println(len(ports))
	fmt.Printf("PortScan1扫描所用时间为：%v 秒\n", time.Since(start))
}

func PortScan1(ip string, ports []int) {
	var wg sync.WaitGroup
	var defaultport int = 65536
	stime := time.Now()
	timeout := 100 * time.Millisecond
	if defaultport > 1000 {
		splits := defaultport / 1000
		var endvalue int
		if defaultport%1000 != 0 {
			endvalue = defaultport % 1000
		}
		for i := 0; i <= splits; i++ {
			var end int
			start := i*1000 + 1
			if i == splits {
				end = start + endvalue
			} else {
				end = start + 1000
			}
			for port := start; port <= end; port++ {
				wg.Add(1)
				go func(p int) {
					defer wg.Done()
					isopen, _ := IsOpen(ip, p, timeout)
					if isopen {
						ports = append(ports, p)
					}
				}(port)
			}
			wg.Wait()
		}
	}
	fmt.Println(ports)
	fmt.Println(len(ports))
	fmt.Printf("PortScan扫描所用时间为：%v 秒\n", time.Since(stime))
}

func PortScan2(ip string, ports []int) {
	var zwg sync.WaitGroup
	var wg sync.WaitGroup
	var defaultport int = 65536
	stime := time.Now()
	// timeout := 100 * time.Millisecond
	timeout := 10 * time.Second
	if defaultport > 1000 {
		splits := defaultport / 1000
		var endvalue int
		if defaultport%1000 != 0 {
			endvalue = defaultport % 1000
		}
		for i := 0; i <= splits; i++ {
			var end int
			zwg.Add(1)
			start := i*1000 + 1
			if i == splits {
				end = start + endvalue
			} else {
				end = start + 1000
			}
			go func() {
				defer zwg.Done()
				for port := start; port < end; port++ {
					wg.Add(1)
					go func(port int) {
						defer wg.Done()
						isopen, _ := IsOpen(ip, port, timeout)
						if isopen {
							fmt.Printf("open port: %v\n", port)
							ports = append(ports, port)
						}
						file.WriteString(fmt.Sprintf("%v\n", port))
					}(port)
				}
				wg.Wait()
			}()
		}
		zwg.Wait()
	}
	fmt.Println(ports)
	fmt.Println(len(ports))
	fmt.Printf("PortScan扫描所用时间为：%v 秒\n", time.Since(stime))
}

func PortScan3(ip string, ports []int) {
	var wg sync.WaitGroup
	var maxsplit int = 1000
	var defaultport int = 65536
	stime := time.Now()
	if defaultport > maxsplit {
		splits := defaultport / maxsplit
		var endvalue int
		if defaultport%maxsplit != 0 {
			endvalue = defaultport % maxsplit
		}
		for i := 0; i <= splits; i++ {
			var end int
			wg.Add(1)
			start := i * maxsplit
			if i == splits {
				end = start + endvalue
			} else {
				end = start + maxsplit - 1
			}
			go func() {
				defer wg.Done()
				Operation(start, end)
			}()
		}
		wg.Wait()
	}
	fmt.Println(ports)
	fmt.Println(len(ports))
	fmt.Printf("PortScan扫描所用时间为：%v 秒\n", time.Since(stime))
}

func Operation(startport, endport int) {
	var wg sync.WaitGroup
	var count int32
	for port := startport; port <= endport; port++ {
		wg.Add(1)
		count++
		go func(port int) {
			defer wg.Done()
			file.WriteString(fmt.Sprintf("%v\n", port))
		}(port)
	}
	fmt.Println(count)
	wg.Wait()
}

func PortScan4(ip string, ports []int) {
	var wg sync.WaitGroup
	var defaultport int = 65536
	stime := time.Now()
	if defaultport > 1000 {
		splits := defaultport / 1000
		var endvalue int
		if defaultport%1000 != 0 {
			endvalue = defaultport % 1000
		}
		for i := 0; i <= splits; i++ {
			wg.Add(1)
			var end int
			start := i*1000 + 1
			if i == splits {
				end = start + endvalue
			} else {
				end = start + 1000
			}
			go func() {
				defer wg.Done()
				for port := start; port < end; port++ {
					wg.Add(1)
					go func(port int) {
						defer wg.Done()
						file.WriteString(fmt.Sprintf("%v\n", port))
					}(port)
				}
				// wg.Wait()
			}()
		}
		wg.Wait()
	}
	fmt.Println(ports)
	fmt.Println(len(ports))
	fmt.Printf("PortScan扫描所用时间为：%v 秒\n", time.Since(stime))
}

type IsOpenScan func(host string, port int, timeout time.Duration) (bool, error)
type ScnaStruct struct {
	Ip   string
	Port int
	IsOpenScan
}

var ScanChan = make(chan IsOpenScan, 1000)
var ScanStruct = make(chan ScnaStruct, 1000)

func PortScan5(ip string, ports []int) {
	timeout := 200 * time.Millisecond
	for port := 1; port <= 65535; port++ {
		scanstruct := ScnaStruct{
			Ip:         ip,
			Port:       port,
			IsOpenScan: IsOpen,
		}
		ScanStruct <- scanstruct
		// ScanChan <- IsOpen
	}
	// go func() {
	// 	scanFunc := <-ScanChan
	// 	port := 80
	// 	scanFunc(ip, port, timeout)
	// }()
	go func() {
		if scanstruct, ok := <-ScanStruct; ok {
			scanstruct.IsOpenScan(scanstruct.Ip, scanstruct.Port, timeout)
		}
	}()
}

func scanPort(ip string, port int, timeout time.Duration, retries int, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	sem <- struct{}{}
	defer func() { <-sem }()

	address := fmt.Sprintf("%s:%d", ip, port)
	for i := 0; i < retries; i++ {
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err == nil {
			conn.Close()
			fmt.Printf("Port %d is open\n", port)
			return
		}
		time.Sleep(1 * time.Second) // 等待一秒再重试
	}
	// fmt.Printf("Port %d is closed or unreachable\n", port)
}


func PortScanCtx() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dialer := &net.Dialer{
		Timeout:   300 * time.Millisecond,
		KeepAlive: 300 * time.Millisecond,
	}
	conn, err := dialer.DialContext(ctx, "tcp", "ip:port")
	if err != nil {
		fmt.Println("端口未开放")
		return
	} else {
		defer conn.Close()
		fmt.Println("端口开放")
	}
}

func ScanPortRun() {
	var wg sync.WaitGroup

	ip := "10.0.35.64"
	start := time.Now()
	// ports := []int{80, 443, 8080, 3306, 5432}
	timeout := 200 * time.Millisecond
	retries := 3
	maxConcurrency := 65535

	sem := make(chan struct{}, maxConcurrency)

	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		go scanPort(ip, port, timeout, retries, &wg, sem)
	}
	wg.Wait()
	fmt.Printf("Scan completed in %v\n", time.Since(start))
}

func main() {
	// ports := []int{}
	// PortScan("10.0.35.64", ports)
	// PortScan1("10.0.35.64", ports)
	// PortScan2("10.0.35.64", ports)
	// PortScan3("10.0.35.64", ports)
	// PortScan4("10.0.35.64", ports)
	ScanPortRun()
}

```

```go
package main

import (
	"fmt"
	"net"
	"time"
)

func checkUDPPort(address string, timeout time.Duration) (bool, time.Duration) {
	var conn net.Conn
	var err error
	start := time.Now()
	conn, err = net.DialTimeout("udp", address, timeout)

	if err != nil {
		// fmt.Println("DialTimeout error:", err)
		return false, time.Since(start)
	}
	defer conn.Close()

	// Send a message to the server
	_, err = conn.Write([]byte("Hello, World!"))
	if err != nil {
		// fmt.Println("Write error:", err)
		return false, time.Since(start)
	}

	// Set a deadline for reading
	conn.SetReadDeadline(time.Now().Add(timeout))

	// Try to read from the server
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		// fmt.Println("Read error:", err)
		return false, time.Since(start)
	}
	return true, time.Since(start)
}

// checkPort 检测端口是否开放
func checkTCPPort(address string, timeout time.Duration) (bool, time.Duration) {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		// fmt.Println(err)
		return false, time.Since(start)
	}
	conn.Close()
	return true, time.Since(start)
}

func main() {
	ports := []int{22, 23, 53, 80, 443, 10080, 3306, 25}
	addr := "192.168.140.3"
	timeout := 1 * time.Second

	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", addr, port)

		// 检测TCP端口
		isOpen, duration := checkTCPPort(address, timeout)
		fmt.Printf("TCP %s open: %t (checked in %s)\n", address, isOpen, duration)
	}
	for _, port := range ports {
		address := fmt.Sprintf("%s:%d", addr, port)

		// 检测UDP端口
		isOpen, duration := checkUDPPort(address, timeout)
		fmt.Printf("UDP %s open: %t (checked in %s)\n", address, isOpen, duration)
	}
}
```