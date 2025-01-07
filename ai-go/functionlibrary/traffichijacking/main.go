/*
@File   : main.go
@Author : pan
@Time   : 2024-04-29 11:02:12
*/
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	device, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Devices found:", device[1].Name)
	handle, err := pcap.OpenLive(device[0].Name, 65536, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetCount := 0
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		packetCount++
		fmt.Println("Packet:", packetCount)
		fmt.Println(packet)
		// TODO: 进行数据包分析
		time.Sleep(1 * time.Second) // 仅用于示例，避免数据包流量过大
	}
}
