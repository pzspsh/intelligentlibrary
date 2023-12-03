/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 02:20:18
*/
package main

import (
	"flag"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

func processPortItem(port string) []string {
	var ports []string
	arr := strings.Split(port, ",")
	for _, p := range arr {
		if strings.Contains(p, "-") {
			ports = append(ports, rangeToArr(p)...)
		} else {
			ports = append(ports, p)
		}
	}
	return ports
}

func rangeToArr(s string) []string {
	if strings.Contains(s, "-") {
		var arr []string
		from, _ := strconv.Atoi(strings.Split(s, "-")[0])
		to, _ := strconv.Atoi(strings.Split(s, "-")[1])
		if from == 0 {
			from = 1
		}
		if to == 0 {
			to = 65535
		}
		for i := from; i <= to; i++ {
			arr = append(arr, strconv.Itoa(i))
		}
		return arr
	} else {
		return []string{s}
	}
}

func scan(ip string, port string, wg *sync.WaitGroup) {
	conn, err := net.DialTimeout("tcp", ip+":"+port, time.Second)
	if err != nil {
		wg.Done()
		return
	}
	wg.Done()
	defer conn.Close()
	log.Println(ip, port, "端口开放！")
}

func main() {
	ip := flag.String("h", "127.0.0.1", "指定主机IP")
	port := flag.String("p", "1-1000", "指定扫描的端口")
	flag.Parse()
	log.Println("扫描的端口为：", *ip, *port)

	//线程同步
	wg := &sync.WaitGroup{}

	for _, p := range processPortItem(*port) {
		wg.Add(1)
		go scan(*ip, p, wg)
	}
	wg.Wait()
}
