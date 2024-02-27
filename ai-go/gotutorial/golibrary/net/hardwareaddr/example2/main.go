/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 16:14:19
*/
package main

import (
	"errors"
	"net"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func GetMACAddress() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		panic(err.Error())
	}

	mac, macerr := "", errors.New("no valid mac address")
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags&net.FlagLoopback) == 0 && strings.Contains(netInterfaces[i].Flags.String(), "broadcast") {
			index := netInterfaces[i].Index

			if isEthernet(index) {
				mac = netInterfaces[i].HardwareAddr.String()
				return mac, nil
			}
		}
	}
	return mac, macerr
}

// 根据网卡接口 Index 判断其是否为 Ethernet 网卡
func isEthernet(ifindex int) bool {
	aas, err := adapterAddresses()
	if err != nil {
		return false
	}
	result := false
	for _, aa := range aas {
		index := aa.IfIndex
		if ifindex == int(index) {
			switch aa.IfType {
			case windows.IF_TYPE_ETHERNET_CSMACD:
				result = true
			}

			if result {
				break
			}
		}
	}
	return result
}

// 从 net/interface_windows.go 中复制过来
func adapterAddresses() ([]*windows.IpAdapterAddresses, error) {
	var b []byte
	l := uint32(15000) // recommended initial size
	for {
		b = make([]byte, l)
		err := windows.GetAdaptersAddresses(syscall.AF_UNSPEC, windows.GAA_FLAG_INCLUDE_PREFIX, 0, (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])), &l)
		if err == nil {
			if l == 0 {
				return nil, nil
			}
			break
		}
		if err.(syscall.Errno) != syscall.ERROR_BUFFER_OVERFLOW {
			return nil, os.NewSyscallError("getadaptersaddresses", err)
		}
		if l <= uint32(len(b)) {
			return nil, os.NewSyscallError("getadaptersaddresses", err)
		}
	}
	var aas []*windows.IpAdapterAddresses
	for aa := (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])); aa != nil; aa = aa.Next {
		aas = append(aas, aa)
	}
	return aas, nil
}

func main() {

}
