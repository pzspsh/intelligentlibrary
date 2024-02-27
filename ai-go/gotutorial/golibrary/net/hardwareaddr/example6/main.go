/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 16:23:48
*/
package main

import (
	"encoding/json"
	"net"
)

type netInterface struct {
	Index        int       `json:"index"`
	MTU          int       `json:"mtu"`
	Name         string    `json:"name"`
	HardwareAddr string    `json:"hardware_addr"`
	Flags        net.Flags `json:"flags"`
}

// 网络接口信息采集
func GetNetInterface() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	var netInterfaces []netInterface
	for _, n := range interfaces {
		temp := netInterface{
			Index:        n.Index,
			MTU:          n.MTU,
			Name:         n.Name,
			HardwareAddr: n.HardwareAddr.String(),
			Flags:        n.Flags,
		}
		netInterfaces = append(netInterfaces, temp)
	}
	netInterfacesJson, err := json.Marshal(netInterfaces)
	if err != nil {
		panic(err)
	}
	WriteFile("NetInterface.json", netInterfacesJson)
}

func WriteFile(writepath string, b []byte) {

}
