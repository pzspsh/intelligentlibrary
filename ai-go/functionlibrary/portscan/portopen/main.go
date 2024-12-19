/*
@File   : main.go
@Author : pan
@Time   : 2024-12-17 11:46:21
*/
package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var (
	file, _ = os.Create("test.txt")
)

func IsOpen(host string, port int, timeout time.Duration) (bool, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", host, port), timeout)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	return true, nil
}

func PortScan(ip string, ports []int) {
	var wg sync.WaitGroup
	timeout := 100 * time.Millisecond
	threads := make(chan bool, 1000)
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
			start := i * 1000
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
	timeout := 100 * time.Millisecond
	if defaultport > 1000 {
		splits := defaultport / 1000
		var endvalue int
		if defaultport%1000 != 0 {
			endvalue = defaultport % 1000
		}
		for i := 0; i <= splits; i++ {
			var end int
			zwg.Add(1)
			start := i * 1000
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

func main() {
	ports := []int{}
	// PortScan("10.0.35.64", ports)
	// PortScan1("10.0.35.64", ports)
	PortScan2("10.0.35.64", ports)
	// PortScan3("10.0.35.64", ports)
	// PortScan4("10.0.35.64", ports)
}
