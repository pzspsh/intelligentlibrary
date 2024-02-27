/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 19:45:49
*/
package main

import (
	"fmt"
	"net"
)

func GetBroadcastAddress() ([]string, error) {
	broadcastAddress := []string{}

	interfaces, err := net.Interfaces() // 获取所有网络接口
	if err != nil {
		return broadcastAddress, err
	}

	for _, face := range interfaces {
		// 选择 已启用的、能广播的、非回环 的接口
		if (face.Flags & (net.FlagUp | net.FlagBroadcast | net.FlagLoopback)) == (net.FlagBroadcast | net.FlagUp) {
			addrs, err := face.Addrs() // 获取该接口下IP地址
			if err != nil {
				return broadcastAddress, err
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok { // 转换成 IPNet { IP Mask } 形式
					if ipnet.IP.To4() != nil { // 只取IPv4的
						var fields net.IP // 用于存放广播地址字段（共4个字段）
						for i := 0; i < 4; i++ {
							fields = append(fields, (ipnet.IP.To4())[i]|(^ipnet.Mask[i])) // 计算广播地址各个字段
						}
						broadcastAddress = append(broadcastAddress, fields.String()) // 转换为字符串形式
					}
				}
			}
		}
	}

	return broadcastAddress, nil
}

func main() {
	addrs, err := GetBroadcastAddress()
	if err != nil {
		fmt.Println(err)
	}
	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
