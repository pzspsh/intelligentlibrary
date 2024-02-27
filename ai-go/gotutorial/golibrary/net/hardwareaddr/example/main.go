/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 16:11:20
*/
package main

import (
	"log"
	"net"
	"strconv"
	"strings"
)

type Nic struct {
	Index        int
	Name         string
	Ipv4         *net.IPNet
	Mac          net.HardwareAddr
	Mtu          int
	Alive        bool
	Broadcast    bool
	Loopback     bool
	PointToPoint bool
	Multicast    bool
}

func parseIp(s string) *net.IPNet {
	i := strings.Index(s, "/")
	if i < 0 {
		log.Print("the ip address should be xxx.xxx.xxx.xxx/24")
		return nil
	}
	addr, mask := s[:i], s[i+1:]
	ip := net.ParseIP(addr)
	n, _ := strconv.Atoi(mask)
	if n > 32 || n < 0 {
		log.Print("ipv4 address' mask should be upper than 0 and lower than 32")
		return nil
	}
	m := net.CIDRMask(n, 8*4)
	return &net.IPNet{IP: ip, Mask: m}
}

func GetInterface() []Nic {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Print(err)
	}
	var nics []Nic
	for i := 0; i < len(interfaces); i++ {
		nic := Nic{
			Index: interfaces[i].Index,
			Name:  interfaces[i].Name,
			Mac:   interfaces[i].HardwareAddr,
			Mtu:   interfaces[i].MTU,
		}
		if interfaces[i].Flags&1 == 1 {
			nic.Alive = true
		}
		if interfaces[i].Flags&2 == 1 {
			nic.Broadcast = true
		}
		if interfaces[i].Flags&4 == 1 {
			nic.Loopback = true
		}
		if interfaces[i].Flags&8 == 1 {
			nic.PointToPoint = true
		}
		if interfaces[i].Flags&16 == 1 {
			nic.Multicast = true
		}
		address, _ := interfaces[i].Addrs()
		for _, v := range address {
			ipStr := v.String()
			if 4 == len(strings.Split(ipStr, ".")) {
				nic.Ipv4 = parseIp(ipStr)
			}
		}
		nics = append(nics, nic)
	}
	return nics
}

func main() {
	interfaces := GetInterface()
	for i := 0; i < len(interfaces); i++ {
		log.Print(interfaces[i])
	}
}
