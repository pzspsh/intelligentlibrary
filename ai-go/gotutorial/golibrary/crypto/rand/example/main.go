/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 11:00:00
*/
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// sessionId函数用来生成一个session ID，即session的唯一标识符
func sessionId() string {
	b := make([]byte, 32)
	//ReadFull从rand.Reader精确地读取len(b)字节数据填充进b
	//rand.Reader是一个全局、共享的密码用强随机数生成器
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	fmt.Println(b)                              //[62 186 123 16 209 19 130 218 146 136 171 211 12 233 45 99 80 200 59 20 56 254 170 110 59 147 223 177 48 136 220 142]
	return base64.URLEncoding.EncodeToString(b) //将生成的随机数b编码后返回字符串,该值则作为session ID
}
func main() {
	fmt.Println(sessionId()) //Prp7ENETgtqSiKvTDOktY1DIOxQ4_qpuO5PfsTCI3I4=
}
