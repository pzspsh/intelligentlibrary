/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 02:01:29
*/
package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var resultPortsList []int
	for i := 1; i <= 50; i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= 1000; i++ {
			ports <- i
		}
	}()
	for i := 0; i < 1000; i++ {
		port := <-results
		if port != 0 {
			resultPortsList = append(resultPortsList, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(resultPortsList)
	for _, port := range resultPortsList {
		fmt.Println(port, "is open")
	}
}
