/*
@File   : main.go
@Author : pan
@Time   : 2024-04-29 15:05:26
*/
package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	device, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	handle, err := pcap.OpenLive(device[0].Name, 65536, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethernetLayer != nil {
			ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
			fmt.Println(ethernetPacket)
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			if ipLayer != nil {
				ipPacket, _ := ipLayer.(*layers.IPv4)
				fmt.Println("Source IP:", ipPacket.SrcIP)
				fmt.Println("Destination IP:", ipPacket.DstIP)

				tcpLayer := packet.Layer(layers.LayerTypeTCP)
				if tcpLayer != nil {
					tcpPacket, _ := tcpLayer.(*layers.TCP)
					fmt.Println("Source Port:", tcpPacket.SrcPort)
					fmt.Println("Destination Port:", tcpPacket.DstPort)
					fmt.Println("Payload:", string(tcpPacket.Payload))
				}

				udpLayer := packet.Layer(layers.LayerTypeUDP)
				if udpLayer != nil {
					udpPacket, _ := udpLayer.(*layers.UDP)
					fmt.Println("Source Port:", udpPacket.SrcPort)
					fmt.Println("Destination Port:", udpPacket.DstPort)
					fmt.Println("Payload:", string(udpPacket.Payload))
				}
			}
		}
	}
}
