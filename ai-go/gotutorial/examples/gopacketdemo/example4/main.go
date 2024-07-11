/*
@File   : main.go
@Author : pan
@Time   : 2024-07-11 13:17:59
*/
package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/examples/util"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
)

var iface = flag.String("i", "eth0", "Interface to get packets from")
var snaplen = flag.Int("s", 1600, "SnapLen for pcap packet capture")

// Build a simple HTTP request parser using tcpassembly.StreamFactory and tcpassembly.Stream interfaces

// httpStreamFactory implements tcpassembly.StreamFactory
type httpStreamFactory struct{}

// httpStream will handle the actual decoding of http requests.
type httpStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
}

func (h *httpStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	hstream := &httpStream{
		net:       net,
		transport: transport,
		r:         tcpreader.NewReaderStream(),
	}
	go hstream.run() // Important... we must guarantee that data from the reader stream is read.

	// ReaderStream implements tcpassembly.Stream, so we can return a pointer to it.
	return &hstream.r
}

func (h *httpStream) run() {
	buf := bufio.NewReader(&h.r)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF {
			// We must read until we see an EOF... very important!
			return
		} else if err != nil {
			log.Println("Error reading stream", h.net, h.transport, ":", err)
		} else {
			bodyBytes := tcpreader.DiscardBytesToEOF(req.Body)
			req.Body.Close()
			log.Println("Received request from stream", h.net, h.transport, ":", req, "with", bodyBytes, "bytes in request body")
		}
	}
}

func main() {
	defer util.Run()()
	var handle *pcap.Handle
	var err error

	// Set up pcap packet capture
	handle, err = pcap.OpenLive(*iface, int32(*snaplen), true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}

	// Set up assembly
	streamFactory := &httpStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	// Read in packets, pass to assembler.
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()
	ticker := time.Tick(time.Minute)
	for {
		select {
		case packet := <-packets:
			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				log.Println("Unusable packet")
				continue
			}
			tcp := packet.TransportLayer().(*layers.TCP)
			//将数据包进行重组
			assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcp, packet.Metadata().Timestamp)

		case <-ticker:
			//每隔一分钟，刷新之前两分钟内不活动的连接
			assembler.FlushOlderThan(time.Now().Add(time.Minute * -2))
		}
	}
}

/*
基本步骤如下：
 1、创建httpStreamFactory结构体，实现tcpassembly.StreamFactory接口
 2、创建连接池
 streamPool := tcpassembly.NewStreamPool(streamFactory)
 3、创建重组器
 assembler := tcpassembly.NewAssembler(streamPool)
 4、将数据包添加到重组器中
 assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcp, packet.Metadata().Timestamp)
*/
