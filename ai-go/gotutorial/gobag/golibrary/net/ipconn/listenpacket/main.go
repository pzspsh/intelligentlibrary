/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 17:35:52
*/
package main

import (
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/examples/util"
	"github.com/google/gopacket/layers"
)

func main() {
	defer util.Run()()

	// XXX create tcp/ip packet
	srcIP := net.ParseIP("127.0.0.1")
	dstIP := net.ParseIP("192.168.0.1")
	//srcIPaddr := net.IPAddr{
	//  IP: srcIP,
	//}
	dstIPaddr := net.IPAddr{
		IP: dstIP,
	}
	ipLayer := layers.IPv4{
		SrcIP:    srcIP,
		DstIP:    dstIP,
		Protocol: layers.IPProtocolTCP,
	}
	tcpLayer := layers.TCP{
		SrcPort: layers.TCPPort(666),
		DstPort: layers.TCPPort(22),
		SYN:     true,
	}
	tcpLayer.SetNetworkLayerForChecksum(&ipLayer)
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	err := gopacket.SerializeLayers(buf, opts, &ipLayer, &tcpLayer)
	if err != nil {
		panic(err)
	}
	// XXX end of packet creation

	// XXX send packet
	ipConn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		panic(err)
	}
	_, err = ipConn.WriteTo(buf.Bytes(), &dstIPaddr)
	if err != nil {
		panic(err)
	}
	log.Print("packet sent!\n")
}
