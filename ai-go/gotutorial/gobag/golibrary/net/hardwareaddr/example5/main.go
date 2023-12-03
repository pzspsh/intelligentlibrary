/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 16:21:38
*/
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"syscall"
	"unsafe"

	"github.com/google/gopacket/pcap"
)

func getAdapterList() (*syscall.IpAdapterInfo, error) {
	b := make([]byte, 1000)
	l := uint32(len(b))
	a := (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
	err := syscall.GetAdaptersInfo(a, &l)
	if err == syscall.ERROR_BUFFER_OVERFLOW {
		b = make([]byte, l)
		a = (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
		err = syscall.GetAdaptersInfo(a, &l)
	}
	if err != nil {
		return nil, os.NewSyscallError("GetAdaptersInfo", err)
	}
	return a, nil
}

func localAddresses() error {
	fmt.Println("测试net.Interfaces()")
	ifaces, err := net.Interfaces()
	if err != nil {
		return err
	}
	for _, iface := range ifaces {
		fmt.Printf("%+v\n", iface)
	}
	fmt.Println()

	fmt.Println("测试net.InterfaceAddrs()")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return err
	}
	for _, addr := range addrs {
		fmt.Printf("\n%+v\n", addr.String())
		fmt.Printf("%+v\n", addr.Network())
	}
	fmt.Println()
	//net.InterfaceByIndex()
	//net.InterfaceByName()
	//for _, addr := range addrs {
	//	fmt.Printf("%T", addr)
	//	// 这个网络地址是IP地址: ipv4, ipv6
	//	ipNet1, isIpNet1 := addr.(*net.IPNet)
	//	fmt.Println(ipNet1)
	//	fmt.Println(isIpNet1)
	//	if ipNet, isIpNet := addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
	//		// 跳过IPV6
	//		if ipNet.IP.To4() != nil {
	//			ipv4 := ipNet.IP.String() // 192.168.1.1
	//			fmt.Println(ipv4)
	//		}
	//	}
	//}

	fmt.Println("测试基于syscall的getAdapterList()")
	aList, err := getAdapterList()
	if err != nil {
		return err
	}
	for ai := aList; ai != nil; ai = ai.Next {
		fmt.Printf("\nIndex:\t%v\n", ai.Index)
		fmt.Printf("AdapterName:\t%s\n", &ai.AdapterName)
		fmt.Printf("Description:\t%s\n", &ai.Description)
		fmt.Printf("Address:\t%s\n", &ai.Address)
		ipl := &ai.IpAddressList
		gwl := &ai.GatewayList
		dhcpl := &ai.DhcpServer
		for ; ipl != nil; ipl = ipl.Next {
			fmt.Printf("IpAddress:\t%s\n", ipl.IpAddress)
			fmt.Printf("IpMask:\t%s\n", ipl.IpMask)
			fmt.Printf("GatewayIp:\t%s\n", gwl.IpAddress)
			fmt.Printf("DHCPServerIp:\t%s\n", dhcpl.IpAddress)
		}
	}
	fmt.Println()

	return err
}

func main() {
	if e := localAddresses(); e != nil {
		fmt.Println(e)
	}

	fmt.Println("测试pcap.FindAllDevs()")
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	// Print device information
	for _, device := range devices {
		fmt.Println("\nName: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Devices Flags: ", device.Flags)
		fmt.Println("Devices addresses: ")
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
			fmt.Println("- Broadaddr: ", address.Broadaddr)
		}
	}
}
