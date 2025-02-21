/*
@File   : main.go
@Author : pan
@Time   : 2025-02-21 17:34:09
*/
package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/routing"
)

func main() {
	targetIP := "ip"  // 目标IP
	targetPort := 445 // 目标端口

	// 获取路由信息以确定网络接口和网关MAC
	router, err := routing.New()
	if err != nil {
		log.Fatal("路由初始化失败:", err)
	}
	iface, gwIP, srcIP, err := router.Route(net.ParseIP(targetIP))
	if err != nil {
		log.Fatal("路由查找失败:", err)
	}

	// 打开网络接口捕获数据包
	handle, err := pcap.OpenLive(iface.Name, 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal("打开网络接口失败:", err)
	}
	defer handle.Close()

	// 发送SYN包
	if err := sendSYN(handle, iface, srcIP, net.ParseIP(targetIP), targetPort, gwIP); err != nil {
		log.Fatal("发送SYN失败:", err)
	}

	// 捕获响应包
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	timeout := time.After(5 * time.Second)
	for {
		select {
		case packet := <-packetSource.Packets():
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			tcpLayer := packet.Layer(layers.LayerTypeTCP)
			if ipLayer == nil || tcpLayer == nil {
				continue
			}

			ip, _ := ipLayer.(*layers.IPv4)
			tcp, _ := tcpLayer.(*layers.TCP)
			if tcp.SYN && tcp.ACK {
				os := detectOS(ip.TTL, tcp.Window)
				fmt.Printf("IP: %s 端口: %d 操作系统可能为: %s\n", targetIP, targetPort, os)
				return
			}
		case <-timeout:
			fmt.Println("超时：未收到响应")
			return
		}
	}
}

// 发送SYN包
func sendSYN(handle *pcap.Handle, iface *net.Interface, srcIP, dstIP net.IP, dstPort int, gwIP net.IP) error {
	// 获取目标MAC地址
	dstMAC, err := getMAC(iface, srcIP, dstIP, gwIP)
	if err != nil {
		return err
	}

	// 构造以太网层
	eth := &layers.Ethernet{
		SrcMAC:       iface.HardwareAddr,
		DstMAC:       dstMAC,
		EthernetType: layers.EthernetTypeIPv4,
	}

	// 构造IP层
	ip := &layers.IPv4{
		SrcIP:    srcIP,
		DstIP:    dstIP,
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolTCP,
	}

	// 构造TCP层
	tcp := &layers.TCP{
		SrcPort: layers.TCPPort(12345), // 随机源端口
		DstPort: layers.TCPPort(dstPort),
		SYN:     true,
		Window:  65535,
	}
	tcp.SetNetworkLayerForChecksum(ip)

	// 序列化包
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}
	if err := gopacket.SerializeLayers(buf, opts, eth, ip, tcp); err != nil {
		return err
	}

	// 发送包
	return handle.WritePacketData(buf.Bytes())
}

// 获取目标MAC地址（通过ARP）
func getMAC(iface *net.Interface, srcIP, dstIP, gwIP net.IP) (net.HardwareAddr, error) {
	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return nil, err
	}
	defer handle.Close()

	// 发送ARP请求
	arpDst := dstIP
	if gwIP != nil {
		arpDst = gwIP
	}
	mac, err := getHwAddr(handle, iface, srcIP, arpDst)
	if err != nil {
		return nil, err
	}
	return mac, nil
}

// 检测操作系统（简化版指纹匹配）
func detectOS(ttl uint8, window uint16) string {
	// TTL匹配
	switch {
	case ttl <= 64:
		return "Linux/Unix"
	case ttl <= 128:
		return "Windows"
	default:
		return "Unknown"
	}

	// 更复杂的指纹可结合窗口大小和TCP选项
	// 例如：Windows默认窗口大小为64240，Linux为29200
}

// 新增到main函数下方
func getHwAddr(handle *pcap.Handle, iface *net.Interface, srcIP, dstIP net.IP) (net.HardwareAddr, error) {
	// 构造ARP请求包
	eth := layers.Ethernet{
		SrcMAC:       iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, // 广播地址
		EthernetType: layers.EthernetTypeARP,
	}
	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPRequest,
		SourceHwAddress:   []byte(iface.HardwareAddr),
		SourceProtAddress: []byte(srcIP.To4()),
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},
		DstProtAddress:    []byte(dstIP.To4()),
	}

	// 序列化数据包
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	if err := gopacket.SerializeLayers(buf, opts, &eth, &arp); err != nil {
		return nil, err
	}

	// 发送ARP请求
	if err := handle.WritePacketData(buf.Bytes()); err != nil {
		return nil, err
	}

	// 设置响应捕获超时
	start := time.Now()
	for time.Since(start) < time.Second*3 {
		// 读取响应数据包
		data, _, err := handle.ReadPacketData()
		if err != nil {
			continue
		}

		// 解析数据包
		packet := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.NoCopy)
		if arpLayer := packet.Layer(layers.LayerTypeARP); arpLayer != nil {
			arpResponse, _ := arpLayer.(*layers.ARP)
			if arpResponse.Operation == layers.ARPReply &&
				net.IP(arpResponse.SourceProtAddress).Equal(dstIP) {
				return net.HardwareAddr(arpResponse.SourceHwAddress), nil
			}
		}
	}
	return nil, fmt.Errorf("ARP请求超时")
}
