/*
@File   : main.go
@Author : pan
@Time   : 2024-04-29 14:04:31
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
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

var iface = flag.String("i", "en0", "Interface to get packets from")
var fname = flag.String("r", "", "Filename to read from, overrides -i")
var snaplen = flag.Int("s", 1600, "SnapLen for pcap packet capture")
var filter = flag.String("f", "tcp and port 80", "BPF filter for pcap")
var logAllPackets = flag.Bool("v", false, "Logs every packet in great detail")

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
	src, _ := transport.Endpoints()
	if fmt.Sprintf("%v", src) == "80" {
		go hstream.runResponse() // Important... we must guarantee that data from the reader stream is read.
	} else {
		go hstream.runRequest() // Important... we must guarantee that data from the reader stream is read.
	}
	return &hstream.r // ReaderStream implements tcpassembly.Stream, so we can return a pointer to it.
}

func (h *httpStream) runResponse() {
	buf := bufio.NewReader(&h.r)
	defer tcpreader.DiscardBytesToEOF(buf)
	for {
		resp, err := http.ReadResponse(buf, nil)
		if err == io.EOF || err == io.ErrUnexpectedEOF { // We must read until we see an EOF... very important!
			return
		} else if err != nil {
			log.Println("Error reading stream", h.net, h.transport, ":", err)
			return
		} else {
			bodyBytes := tcpreader.DiscardBytesToEOF(resp.Body)
			resp.Body.Close()
			printResponse(resp, h, bodyBytes)
			log.Println("Received response from stream", h.net, h.transport, ":", resp, "with", bodyBytes, "bytes in response body")
		}
	}
}

func (h *httpStream) runRequest() {
	buf := bufio.NewReader(&h.r)
	defer tcpreader.DiscardBytesToEOF(buf)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF || err == io.ErrUnexpectedEOF { // We must read until we see an EOF... very important!
			return
		} else if err != nil {
			log.Println("Error reading stream", h.net, h.transport, ":", err)
		} else {
			bodyBytes := tcpreader.DiscardBytesToEOF(req.Body)
			req.Body.Close()
			printRequest(req, h, bodyBytes)
			// log.Println("Received request from stream", h.net, h.transport, ":", req, "with", bodyBytes, "bytes in request body")
		}
	}
}

func printHeader(h http.Header) {
	for k, v := range h {
		fmt.Println(k, v)
	}
}

func printRequest(req *http.Request, h *httpStream, bodyBytes int) {
	fmt.Println("\n\r\n\r")
	fmt.Println(h.net, h.transport, bodyBytes)
	fmt.Println("\n\r")
	fmt.Println(req.Method, req.RequestURI, req.Proto)
	printHeader(req.Header)

}

func printResponse(resp *http.Response, h *httpStream, bodyBytes int) {
	fmt.Println("\n\r")
	fmt.Println(h.net, h.transport, bodyBytes)
	fmt.Println(resp.Proto, resp.Status)
	printHeader(resp.Header)
}

func main() {
	defer util.Run()()
	var handle *pcap.Handle
	var err error

	device, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	if *fname != "" { // Set up pcap packet capture
		log.Printf("Reading from pcap dump %q", *fname)
		handle, err = pcap.OpenOffline(*fname)
	} else {
		log.Printf("Starting capture on interface %q", *iface)
		// handle, err = pcap.OpenLive(*iface, int32(*snaplen), true, pcap.BlockForever)
		log.Printf("Starting capture on interface %q", device[0].Name)
		handle, err = pcap.OpenLive(device[0].Name, int32(*snaplen), true, pcap.BlockForever)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err := handle.SetBPFFilter(*filter); err != nil {
		log.Fatal(err)
	}

	// Set up assembly
	streamFactory := &httpStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	log.Println("reading in packets")
	// Read in packets, pass to assembler.
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()
	ticker := time.Tick(time.Minute)
	for {
		select {
		case packet := <-packets:
			if packet == nil { // A nil packet indicates the end of a pcap file.
				return
			}
			if *logAllPackets {
				log.Println(packet)
			}
			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				log.Println("Unusable packet")
				continue
			}
			tcp := packet.TransportLayer().(*layers.TCP)
			assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcp, packet.Metadata().Timestamp)
		case <-ticker:
			// Every minute, flush connections that haven't seen activity in the past 2 minutes.
			assembler.FlushOlderThan(time.Now().Add(time.Minute * -2))
		}
	}
}
