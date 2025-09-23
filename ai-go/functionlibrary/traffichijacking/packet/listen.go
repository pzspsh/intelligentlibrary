package packet

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
)

// Listen 监听网卡.
func (slf *Handle) Listen() error {
	// 获取网卡信息
	addrs, _ := net.Interfaces()
	// fmt.Println("addrs: ", addrs[1].Name)
	iface, err := net.InterfaceByName(addrs[1].Name)
	if err != nil {
		return fmt.Errorf("cardName %s not found, err: %v", slf.cardName, err)
	}
	log.Printf("cardName: %s, MTU: %d", slf.cardName, iface.MTU)

	// 打开设备监听
	device := "\\Device\\NPF_{B7E01C74-876F-44B9-95FF-13BAB4AB8D9E}"
	handle, err := pcap.OpenLive(device, 1024*1024, slf.promisc, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("openLive %s err: %v", slf.cardName, err)
	}
	defer handle.Close()

	// 设置过滤器
	if err := handle.SetBPFFilter(slf.bpf); err != nil {
		return fmt.Errorf("set bpf filter: %v", err)
	}

	go slf.EventHandle()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	streamFactory := NewHTTPStreamFactory(slf.eventCh)
	pool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(pool)
	ticker := time.NewTicker(time.Minute)
	var lastPacketTimestamp time.Time
	for {
		select {
		case <-slf.ctx.Done():
			return nil
		case packet := <-packetSource.Packets():
			netLayer := packet.NetworkLayer()
			if netLayer == nil {
				continue
			}
			transLayer := packet.TransportLayer()
			if transLayer == nil {
				continue
			}
			tcp, _ := transLayer.(*layers.TCP)
			if tcp == nil {
				continue
			}
			assembler.AssembleWithTimestamp(
				netLayer.NetworkFlow(),
				tcp,
				packet.Metadata().CaptureInfo.Timestamp)

			lastPacketTimestamp = packet.Metadata().CaptureInfo.Timestamp
		case <-ticker.C:
			assembler.FlushOlderThan(lastPacketTimestamp.Add(slf.flushTime))
		}
	}
}
