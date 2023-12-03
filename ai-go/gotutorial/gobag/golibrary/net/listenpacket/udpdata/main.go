/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:23:41
*/
package main

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func EncodeUDPPacket(localIP, remoteIP net.IP, localPort, remotePort uint16, payload []byte) ([]byte, error) {
	ip := &layers.IPv4{
		Version:  4,
		TTL:      128,
		SrcIP:    localIP,
		DstIP:    remoteIP,
		Protocol: layers.IPProtocolUDP,
	}
	udp := &layers.UDP{
		SrcPort: layers.UDPPort(localPort),
		DstPort: layers.UDPPort(remotePort),
	}
	udp.SetNetworkLayerForChecksum(ip)
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}
	err := gopacket.SerializeLayers(buf, opts, udp, gopacket.Payload(payload))
	return buf.Bytes(), err
}
