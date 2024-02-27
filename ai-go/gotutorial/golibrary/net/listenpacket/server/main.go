/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:23:41
*/
package main

// import (
// 	"flag"
// 	"log"
// 	"net"

// 	"github.com/google/gopacket"
// 	"github.com/google/gopacket/layers"
// 	"github.com/smallnest/rpcx/codec"
// 	"golang.org/x/net/bpf"
// 	"golang.org/x/net/ipv4"
// )

// var (
// 	addr = flag.String("s", "localhost", "server address")
// 	port = flag.Int("p", 8972, "port")
// )
// var (
// 	stat         = make(map[string]int)
// 	lastStatTime = int64(0)
// )

// func main() {
// 	flag.Parse()
// 	conn, err := net.ListenPacket("ip4:udp", *addr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	cc := conn.(*net.IPConn)
// 	cc.SetReadBuffer(20 * 1024 * 1024)
// 	cc.SetWriteBuffer(20 * 1024 * 1024)
// 	pconn := ipv4.NewPacketConn(conn)
// 	var assembled []bpf.RawInstruction
// 	if assembled, err = bpf.Assemble(filter); err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	pconn.SetBPF(assembled)
// 	handleConn(conn)
// }
// func handleConn(conn net.PacketConn) {
// 	for {
// 		buffer := make([]byte, 1024)
// 		n, remoteaddr, err := conn.ReadFrom(buffer)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		buffer = buffer[:n]
// 		packet := gopacket.NewPacket(buffer, layers.LayerTypeUDP, gopacket.NoCopy)
// 		// Get the UDP layer from this packet
// 		if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
// 			udp, _ := udpLayer.(*layers.UDP)
// 			if app := packet.ApplicationLayer(); app != nil {
// 				data, err := codec.EncodeUDPPacket(net.ParseIP("127.0.0.1"), net.ParseIP("127.0.0.1"), uint16(udp.DstPort), uint16(udp.SrcPort), app.Payload())
// 				if err != nil {
// 					log.Printf("failed to EncodePacket: %v", err)
// 					return
// 				}
// 				if _, err := conn.WriteTo(data, remoteaddr); err != nil {
// 					log.Printf("failed to write packet: %v", err)
// 					conn.Close()
// 					return
// 				}
// 			}
// 		}
// 	}
// }

// type Filter []bpf.Instruction

// var filter = Filter{
// 	bpf.LoadAbsolute{Off: 22, Size: 2},                                // load the destination port
// 	bpf.JumpIf{Cond: bpf.JumpEqual, Val: uint32(*port), SkipFalse: 1}, // if Val != 8972 skip next instruction
// 	bpf.RetConstant{Val: 0xffff},                                      // return 0xffff bytes (or less) from packet
// 	bpf.RetConstant{Val: 0x0},                                         // return 0 bytes, effectively ignore this packet
// }
