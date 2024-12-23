/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 01:13:19
*/
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Protocol struct {
	Format []string
}

func main() {
	str := "693a0210afc541528457ac135a3fca38"
	p := &Protocol{}
	s2 := p.HexStringToByte(str)
	s1 := p.DecToHexString(s2)
	fmt.Println(s1, string(s2), len(s2))
}

// 编码
func (p *Protocol) Pack(args ...int64) (ret []byte) {
	la := len(args)
	ls := len(p.Format)
	if ls > 0 && la > 0 && ls == la {
		for i := 0; i < la; i++ {
			if p.Format[i] == "N8" {
				ret = append(ret, IntToBytes8(args[i])...)
			} else if p.Format[i] == "N4" {
				ret = append(ret, IntToBytes4(args[i])...)
			} else if p.Format[i] == "N2" {
				ret = append(ret, IntToBytes2(args[i])...)
			}
		}
	}
	return ret
}

// 解码
func (p *Protocol) UnPack(data []byte) []int64 {
	la := len(p.Format)
	ret := make([]int64, la)
	if la > 0 {
		for i := 0; i < la; i++ {
			if p.Format[i] == "N8" {
				ret[i] = Bytes8ToInt64(data[0:8])
				data = data[8:]
			} else if p.Format[i] == "N4" {
				ret[i] = Bytes4ToInt64(data[0:4])
				data = data[4:]
			} else if p.Format[i] == "N2" {
				ret[i] = Bytes2ToInt64(data[0:2])
				data = data[2:]
			}
		}
	}
	return ret
}

// 转成16进制编码字符串
func (p *Protocol) Pack16(args ...int64) (hString string) {
	//变成 byte 码
	hByte := p.Pack(args...)
	//转成 16进制字符串
	hString = p.DecToHexString(hByte)
	return hString
}

// 解码16进制字符串
func (p *Protocol) UnPack16(hString string) (unIntList []int64) {
	HSByte := p.HexStringToByte(hString)
	unIntList = p.UnPack(HSByte)
	return unIntList
}

// 10进制转16进制字符串
func (p *Protocol) DecToHexString(decString []byte) (responseStr string) {
	for _, v := range decString {
		hexString := DecHex(int64(v))
		if len(hexString) < 2 {
			hexString = "0" + hexString
		}
		responseStr += hexString
	}
	return strings.ToLower(responseStr)
}

// 16进制字符串转字节类型
func (p *Protocol) HexStringToByte(hexString string) (responseByte []byte) {
	ls := len(hexString) / 2
	for i := 0; i < ls; i++ {
		hex := hexString[0:2]
		hexString = hexString[2:]
		responseByte = append(responseByte, IntToBytes1(HexDec(hex))...)
	}
	return
}

// int64 转 byte8
func IntToBytes8(n int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(n))
	return buf
}

// int64 转 byte4
func IntToBytes4(n int64) []byte {
	nb := intToBytes(n, 4)
	return nb
}

// int64 转 byte2
func IntToBytes2(n int64) []byte {
	nb := intToBytes(n, 2)
	return nb
}

// int64 转 byte1
func IntToBytes1(n int64) []byte {
	nb := intToBytes(n, 1)
	return nb
}

// int64 转 byteN
func intToBytes(n int64, k int) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	gbyte := bytesBuffer.Bytes()
	//c++ 高低位转换
	x := len(gbyte)
	nb := make([]byte, k)
	for i := 0; i < k; i++ {
		nb[i] = gbyte[x-i-1]
	}
	return nb
}

// byte2 转 int64
func Bytes2ToInt64(b []byte) int64 {
	nb := []byte{0, 0, b[1], b[0]}
	bytesBuffer := bytes.NewBuffer(nb)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int64(x)
}

// byte4 转 int64
func Bytes4ToInt64(b []byte) int64 {
	nb := []byte{b[3], b[2], b[1], b[0]}
	bytesBuffer := bytes.NewBuffer(nb)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int64(x)
}

// byte8 转 int64
func Bytes8ToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// 10进制 转 16进制
func DecHex(n int64) string {
	if n < 0 {
		log.Println("Decimal to hexadecimal error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			m = hex[m]
			s = fmt.Sprintf("%v%v", string(m), s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}

// 16进制 转 10进制
func HexDec(h string) (n int64) {
	s := strings.Split(strings.ToUpper(h), "")
	l := len(s)
	i := 0
	d := float64(0)
	hex := map[string]string{"A": "10", "B": "11", "C": "12", "D": "13", "E": "14", "F": "15"}
	for i = 0; i < l; i++ {
		c := s[i]
		if v, ok := hex[c]; ok {
			c = v
		}
		f, err := strconv.ParseFloat(c, 10)
		if err != nil {
			log.Println("Hexadecimal to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(16, float64(l-i-1))
	}
	return int64(d)
}
