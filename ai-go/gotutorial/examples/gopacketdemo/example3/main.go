/*
@File   : main.go
@Author : pan
@Time   : 2024-07-11 13:17:14
*/
package main

import (
	"fmt"

	"github.com/google/gopacket"
)

// 创建自定义层数据结构，并实现Layer接口中的函数LayerType()、LayerContents()、LayerPayload()
type CustomLayer struct {
	// This layer just has two bytes at the front
	SomeByte    byte
	AnotherByte byte
	restOfData  []byte
}

// 注册自定义层类型，然后我们才可以使用它
// 第一个参数是ID. 自定义层使用大于2000的数字，它必须是唯一的
var CustomLayerType = gopacket.RegisterLayerType(2001, gopacket.LayerTypeMetadata{"CustomLayerType", gopacket.DecodeFunc(decodeCustomLayer)})

// 自定义层实现LayerType
func (l CustomLayer) LayerType() gopacket.LayerType {
	return CustomLayerType
}

// 自定义层实现LayerContents
func (l CustomLayer) LayerContents() []byte {
	return []byte{l.SomeByte, l.AnotherByte}
}

// 自定义层实现LayerPayload
func (l CustomLayer) LayerPayload() []byte {
	return l.restOfData
}

// 实现自定义的解码函数
func decodeCustomLayer(data []byte, p gopacket.PacketBuilder) error {
	p.AddLayer(&CustomLayer{data[0], data[1], data[2:]})
	return p.NextDecoder(gopacket.LayerTypePayload)
}

func main() {
	rawBytes := []byte{0xF0, 0x0F, 65, 65, 66, 67, 68}
	packet := gopacket.NewPacket(
		rawBytes,
		CustomLayerType,
		gopacket.Default,
	)
	fmt.Println("Created packet out of raw bytes.")
	fmt.Println(packet)
	// Decode the packet as our custom layer
	customLayer := packet.Layer(CustomLayerType)
	if customLayer != nil {
		fmt.Println("Packet was successfully decoded with custom layer decoder.")
		customLayerContent, _ := customLayer.(*CustomLayer)
		// Now we can access the elements of the custom struct
		fmt.Println("Payload: ", customLayerContent.LayerPayload())
		fmt.Println("SomeByte element:", customLayerContent.SomeByte)
		fmt.Println("AnotherByte element:", customLayerContent.AnotherByte)
	}
}
