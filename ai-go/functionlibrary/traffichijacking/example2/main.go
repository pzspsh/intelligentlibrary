/*
@File   : main.go
@Author : pan
@Time   : 2024-04-29 15:02:34
*/
package main

/* import (
	"fmt"
	"log"
	"strings"
	"time"

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
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			if ipLayer != nil {
				ipPacket, _ := ipLayer.(*layers.IPv4)
				tcpLayer := packet.Layer(layers.LayerTypeTCP)
				if tcpLayer != nil {
					tcpPacket, _ := tcpLayer.(*layers.TCP)
					httpLayer := packet.Layer(layers.LayerTypeHTTP)
					if httpLayer != nil {
						httpPacket, _ := httpLayer.(*layers.HTTP)
						fmt.Println("Source MAC:", ethernetPacket.SrcMAC)
						fmt.Println("Destination MAC:", ethernetPacket.DstMAC)

						fmt.Println("Source IP:", ipPacket.SrcIP)
						fmt.Println("Destination IP:", ipPacket.DstIP)

						fmt.Println("Source Port:", tcpPacket.SrcPort)
						fmt.Println("Destination Port:", tcpPacket.DstPort)

						fmt.Println("HTTP Method:", httpPacket.Method)
						fmt.Println("HTTP Host:", httpPacket.Host)

						headers := strings.Split(string(httpPacket.Headers), "\r\n")
						for _, header := range headers {
							fmt.Println("HTTP Header:", header)
						}

						fmt.Println("--------")
					}
				}
			}
		}

		time.Sleep(1 * time.Second) // 仅用于示例，避免数据包流量过大
	}
} */

