/*
@File   : main.go
@Author : pan
@Time   : 2024-12-17 11:46:21
*/
package main

import (
	"fmt"
)

type OpenPort struct {
	Ip     string
	Port   int
	Banner string
}

func IsOpen(ip string, port int, retries int) (bool, error) {
	var err error
	var isopen bool

	return isopen, err
}

func PortScan(ip string, ports []int, retrie, threads int) {

}

func Param() {

}

func main() {
	var threads int = 1000
	var retries int = 3
	var ports []int = []int{22, 80, 443}
	fmt.Println(threads, retries, ports)
}
