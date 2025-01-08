/*
@File   : main.go
@Author : pan
@Time   : 2024-04-29 11:02:12
*/
package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/google/gopacket/pcap"
// )

// /*
// 先调用pcap.FindAllDevs()获取当前主机所有的网络设备，网络设备有哪些属性
// */
// func main() {
// 	// 得到所有的(网络)设备
// 	devices, err := pcap.FindAllDevs()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// 打印设备信息
// 	fmt.Println("Devices found:")
// 	for _, device := range devices {
// 		fmt.Println("\nName: ", device.Name)
// 		fmt.Println("Description: ", device.Description)
// 		fmt.Println("Devices addresses: ", device.Description)
// 		for _, address := range device.Addresses {
// 			fmt.Println("- IP address: ", address.IP)
// 			fmt.Println("- Subnet mask: ", address.Netmask)
// 		}
// 	}
// }

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func GetPcapInfo() {
	// 得到所有的(网络)设备
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	// 打印设备信息
	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("\nName: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Devices addresses: ", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
}

func PacketDemo() {
	// device, err := pcap.FindAllDevs()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	device := "\\Device\\NPF_{B7E01C74-876F-44B9-95FF-13BAB4AB8D9E}"
	handle, err := pcap.OpenLive(device, 65536, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethernetLayer != nil {
			ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
			fmt.Println("Source MAC:", ethernetPacket.SrcMAC)
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

func printPacketInfo(packet gopacket.Packet) {
	// Let's see if the packet is an ethernet packet
	// 判断数据包是否为以太网数据包，可解析出源mac地址、目的mac地址、以太网类型（如ip类型）等
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		fmt.Println("Ethernet layer detected.")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
		// Ethernet type is typically IPv4 but could be ARP or other
		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
		fmt.Println()
	}
	// Let's see if the packet is IP (even though the ether type told us)
	// 判断数据包是否为IP数据包，可解析出源ip、目的ip、协议号等
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 layer detected.")
		ip, _ := ipLayer.(*layers.IPv4)
		// IP layer variables:
		// Version (Either 4 or 6)
		// IHL (IP Header Length in 32-bit words)
		// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
		// Checksum, SrcIP, DstIP
		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("Protocol: ", ip.Protocol)
		fmt.Println()
	}
	// Let's see if the packet is TCP
	// 判断数据包是否为TCP数据包，可解析源端口、目的端口、seq序列号、tcp标志位等
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		fmt.Println("TCP layer detected.")
		tcp, _ := tcpLayer.(*layers.TCP)
		// TCP layer variables:
		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
		fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequence number: ", tcp.Seq)
		fmt.Println()
	}
	// Iterate over all layers, printing out each layer type
	fmt.Println("All packet layers:")
	for _, layer := range packet.Layers() {
		fmt.Println("- ", layer.LayerType())
	}
	///.......................................................
	// Check for errors
	// 判断layer是否存在错误
	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet:", err)
	}
}

var (
	device      string = "/Device/NPF_{B7E01C74-876F-44B9-95FF-13BAB4AB8D9E}"
	snapshotLen int32  = 1024
	promiscuous bool   = false
	err         error
	timeout     time.Duration = 30 * time.Second
	handle      *pcap.Handle
)

func PacketInfoRun() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		printPacketInfo(packet)
	}
}

/* 网络嗅探器 */
func main() {
	// GetPcapInfo()
	PacketInfoRun()
}
