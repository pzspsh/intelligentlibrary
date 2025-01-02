/*
@File   : main.go
@Author : pan
@Time   : 2025-01-02 15:35:54
*/
package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	// 得到所有的(网络)设备
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	// 打印设备信息
	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("Name: ", device.Name)
		fmt.Println("Description: ", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("- IP地址: ", address.IP)
			fmt.Println("- 子网掩码: ", address.Netmask)
		}
		fmt.Println("")
	}
}