/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 15:19:18
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"mime"
)

func main() {
	dec := new(mime.WordDecoder)
	header, err := dec.DecodeHeader("=?utf-8?q?=C3=89ric?= <eric@example.org>, =?utf-8?q?Ana=C3=AFs?= <anais@example.org>")
	if err != nil {
		panic(err)
	}
	fmt.Println(header)

	header, err = dec.DecodeHeader("=?utf-8?q?=C2=A1Hola,?= =?utf-8?q?_se=C3=B1or!?=")
	if err != nil {
		panic(err)
	}
	fmt.Println(header)

	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "x-case":
			// Fake character set for example.
			// Real use would integrate with packages such
			// as code.google.com/p/go-charset
			content, err := io.ReadAll(input)
			if err != nil {
				return nil, err
			}
			return bytes.NewReader(bytes.ToUpper(content)), nil
		default:
			return nil, fmt.Errorf("unhandled charset %q", charset)
		}
	}
	header, err = dec.DecodeHeader("=?x-case?q?hello_?= =?x-case?q?world!?=")
	if err != nil {
		panic(err)
	}
	fmt.Println(header)
}
