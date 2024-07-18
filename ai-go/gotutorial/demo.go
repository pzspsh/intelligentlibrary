/*
@File   : demo.go
@Author : pan
@Time   : 2024-07-18 15:37:16
*/
package main

import (
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/rs/xid"
	"gopkg.in/corvus-ch/zbase32.v1"
)

func Url(host string) string {
	data := make([]byte, 13)
	correlationID := xid.New().String()
	_, _ = rand.Read(data)
	randomData := zbase32.StdEncoding.EncodeToString(data)
	if len(randomData) > 13 {
		randomData = randomData[:13]
	}
	builder := &strings.Builder{}
	builder.Grow(len(correlationID) + len(randomData) + len(host) + 1)
	builder.WriteString(correlationID)
	builder.WriteString(randomData)
	builder.WriteString(".")
	builder.WriteString(host)
	URL := builder.String()
	return URL
}

func main() {
	url := Url("baidu.com")
	fmt.Println(url)
}
